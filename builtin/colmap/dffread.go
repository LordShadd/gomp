package colmap

/*
#include "lib/bullet_capi.h"
*/
import "C"
import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

const (
	rwChunkStruct       = uint32(0x01)
	rwChunkExtension    = uint32(0x03)
	rwChunkFrameList    = uint32(0x0E)
	rwChunkGeometryList = uint32(0x1A)
	rwChunkAtomic       = uint32(0x14)
	rwChunkClump        = uint32(0x10)
	rwChunkSampCol      = uint32(0x253F2FF)
)

type rwHeader struct {
	Type, Length, Version uint32
}

func readRWHeader(r io.Reader) (rwHeader, error) {
	var h rwHeader
	return h, binary.Read(r, binary.LittleEndian, &h)
}

func skipRWChunk(r io.ReadSeeker, length uint32) error {
	_, err := r.Seek(int64(length), io.SeekCurrent)
	return err
}

type colItemsHeader struct {
	NumSpheres   uint16
	NumBoxes     uint16
	NumFaces     uint16
	NumWheels    uint8
	Padding      uint8
	Flags        uint32
	OffSpheres   uint32
	OffBoxes     uint32
	OffSuspLines uint32
	OffVertices  uint32
	OffFaces     uint32
	OffTriPlanes uint32
	NumShadow    uint32
	OffShadVert  uint32
	OffShadFaces uint32
}

func LoadFromDff(modelID int, dffPath string) bool {
	f, err := os.Open(dffPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "colmap: LoadFromDff: cannot open %s: %v\n", dffPath, err)
		return false
	}
	defer f.Close()

	h, err := readRWHeader(f)
	if err != nil || h.Type != rwChunkClump {
		fmt.Fprintf(os.Stderr, "colmap: LoadFromDff: not a valid DFF (clump chunk not found)\n")
		return false
	}

	sh, err := readRWHeader(f)
	if err != nil || sh.Type != rwChunkStruct {
		return false
	}
	var numAtomics uint32
	binary.Read(f, binary.LittleEndian, &numAtomics)
	if sh.Length == 0xC {
		f.Seek(8, io.SeekCurrent)
	}

	fh, _ := readRWHeader(f)
	if fh.Type == rwChunkFrameList {
		f.Seek(int64(fh.Length), io.SeekCurrent)
	}

	gh, _ := readRWHeader(f)
	if gh.Type == rwChunkGeometryList {
		f.Seek(int64(gh.Length), io.SeekCurrent)
	}

	ah, _ := readRWHeader(f)
	if ah.Type == rwChunkAtomic {
		f.Seek(int64(ah.Length), io.SeekCurrent)
	}

	eh, err := readRWHeader(f)
	if err != nil || eh.Type != rwChunkExtension {
		return false
	}

	extEnd, _ := f.Seek(0, io.SeekCurrent)
	extEnd += int64(eh.Length)

	for {
		pos, _ := f.Seek(0, io.SeekCurrent)
		if pos >= extEnd {
			break
		}
		ch, err := readRWHeader(f)
		if err != nil {
			break
		}
		if ch.Type == rwChunkSampCol {
			s := parseDFFCollision(f, modelID)
			if s != nil {
				shapeMu.Lock()
				modelShapes[int32(modelID)] = s
				shapeMu.Unlock()
				return true
			}
			return false
		}
		f.Seek(int64(ch.Length), io.SeekCurrent)
	}

	fmt.Fprintf(os.Stderr, "colmap: LoadFromDff: no SAMP collision chunk found in %s\n", dffPath)
	return false
}

func parseDFFCollision(r io.ReadSeeker, modelID int) C.CMShape {
	beginPos, _ := r.Seek(0, io.SeekCurrent)

	var validator [4]byte
	r.Read(validator[:])
	colVersion := validator[3]

	var size uint32
	binary.Read(r, binary.LittleEndian, &size)

	r.Seek(24, io.SeekCurrent)

	s := C.cm_shape_create()

	if colVersion == 'L' {
		if !parseCOL1(r, s) {
			C.cm_shape_destroy(s)
			return nil
		}
		r.Seek(beginPos+8+int64(size), io.SeekStart)
	} else {
		r.Seek(40, io.SeekCurrent)
		if !parseCOL2(r, beginPos, s) {
			C.cm_shape_destroy(s)
			return nil
		}
		r.Seek(beginPos+8+int64(size), io.SeekStart)
	}
	return s
}

func parseCOL1(r io.ReadSeeker, s C.CMShape) bool {
	var count uint32
	binary.Read(r, binary.LittleEndian, &count)
	for i := uint32(0); i < count; i++ {
		var radius float32
		var ox, oy, oz float32
		binary.Read(r, binary.LittleEndian, &radius)
		binary.Read(r, binary.LittleEndian, &ox)
		binary.Read(r, binary.LittleEndian, &oy)
		binary.Read(r, binary.LittleEndian, &oz)
		r.Seek(4, io.SeekCurrent)
		C.cm_shape_add_sphere(s, C.float(ox), C.float(oy), C.float(oz), C.float(radius))
	}
	r.Seek(4, io.SeekCurrent)

	binary.Read(r, binary.LittleEndian, &count)
	for i := uint32(0); i < count; i++ {
		var minX, minY, minZ, maxX, maxY, maxZ float32
		binary.Read(r, binary.LittleEndian, &minX)
		binary.Read(r, binary.LittleEndian, &minY)
		binary.Read(r, binary.LittleEndian, &minZ)
		binary.Read(r, binary.LittleEndian, &maxX)
		binary.Read(r, binary.LittleEndian, &maxY)
		binary.Read(r, binary.LittleEndian, &maxZ)
		r.Seek(4, io.SeekCurrent)
		cx := (minX + maxX) / 2
		cy := (minY + maxY) / 2
		cz := (minZ + maxZ) / 2
		sx := (maxX - minX) / 2
		sy := (maxY - minY) / 2
		sz := (maxZ - minZ) / 2
		C.cm_shape_add_box(s, C.float(cx), C.float(cy), C.float(cz), C.float(sx), C.float(sy), C.float(sz))
	}

	binary.Read(r, binary.LittleEndian, &count)
	if count == 0 {
		return true
	}
	verts := make([][3]float32, count)
	for i := range verts {
		binary.Read(r, binary.LittleEndian, &verts[i][0])
		binary.Read(r, binary.LittleEndian, &verts[i][1])
		binary.Read(r, binary.LittleEndian, &verts[i][2])
	}

	var faceCount uint32
	binary.Read(r, binary.LittleEndian, &faceCount)
	if faceCount == 0 {
		return true
	}
	return addTrimeshFromIndexed(r, s, verts, int(faceCount), 4, false)
}

func parseCOL2(r io.ReadSeeker, beginPos int64, s C.CMShape) bool {
	var items colItemsHeader
	if err := binary.Read(r, binary.LittleEndian, &items); err != nil {
		return false
	}

	if items.NumBoxes > 0 {
		r.Seek(beginPos+4+int64(items.OffBoxes), io.SeekStart)
		for i := uint16(0); i < items.NumBoxes; i++ {
			var minX, minY, minZ, maxX, maxY, maxZ float32
			binary.Read(r, binary.LittleEndian, &minX)
			binary.Read(r, binary.LittleEndian, &minY)
			binary.Read(r, binary.LittleEndian, &minZ)
			binary.Read(r, binary.LittleEndian, &maxX)
			binary.Read(r, binary.LittleEndian, &maxY)
			binary.Read(r, binary.LittleEndian, &maxZ)
			r.Seek(4, io.SeekCurrent)
			cx := (minX + maxX) / 2
			cy := (minY + maxY) / 2
			cz := (minZ + maxZ) / 2
			sx := (maxX - minX) / 2
			sy := (maxY - minY) / 2
			sz := (maxZ - minZ) / 2
			C.cm_shape_add_box(s, C.float(cx), C.float(cy), C.float(cz), C.float(sx), C.float(sy), C.float(sz))
		}
	}

	if items.NumSpheres > 0 {
		r.Seek(beginPos+4+int64(items.OffSpheres), io.SeekStart)
		for i := uint16(0); i < items.NumSpheres; i++ {
			var radius, ox, oy, oz float32
			binary.Read(r, binary.LittleEndian, &radius)
			binary.Read(r, binary.LittleEndian, &ox)
			binary.Read(r, binary.LittleEndian, &oy)
			binary.Read(r, binary.LittleEndian, &oz)
			r.Seek(4, io.SeekCurrent)
			C.cm_shape_add_sphere(s, C.float(ox), C.float(oy), C.float(oz), C.float(radius))
		}
	}

	if items.NumFaces > 0 {
		r.Seek(beginPos+4+int64(items.OffFaces), io.SeekStart)
		faceIdx := make([][3]uint16, items.NumFaces)
		topIdx := uint16(0)
		for i := range faceIdx {
			binary.Read(r, binary.LittleEndian, &faceIdx[i][0])
			binary.Read(r, binary.LittleEndian, &faceIdx[i][1])
			binary.Read(r, binary.LittleEndian, &faceIdx[i][2])
			r.Seek(2, io.SeekCurrent)
			for _, v := range faceIdx[i] {
				if v > topIdx {
					topIdx = v
				}
			}
		}

		r.Seek(beginPos+4+int64(items.OffVertices), io.SeekStart)
		verts := make([][3]float32, topIdx+1)
		for i := range verts {
			var v [3]int16
			binary.Read(r, binary.LittleEndian, &v)
			verts[i][0] = float32(v[0]) / 128.0
			verts[i][1] = float32(v[1]) / 128.0
			verts[i][2] = float32(v[2]) / 128.0
		}

		flatVerts := make([]C.float, len(verts)*3)
		for i, v := range verts {
			flatVerts[i*3] = C.float(v[0])
			flatVerts[i*3+1] = C.float(v[1])
			flatVerts[i*3+2] = C.float(v[2])
		}
		flatIdx := make([]C.int, len(faceIdx)*3)
		for i, fi := range faceIdx {
			flatIdx[i*3] = C.int(fi[0])
			flatIdx[i*3+1] = C.int(fi[1])
			flatIdx[i*3+2] = C.int(fi[2])
		}
		nv := C.int(len(verts))
		nt := C.int(len(faceIdx))
		C.cm_shape_add_trimesh(s, &flatVerts[0], nv, &flatIdx[0], nt, 0)
		C.cm_shape_add_trimesh(s, &flatVerts[0], nv, &flatIdx[0], nt, 1)
	}
	return true
}

func addTrimeshFromIndexed(r io.ReadSeeker, s C.CMShape, verts [][3]float32, faceCount int, skipBytes int64, _ bool) bool {
	flatVerts := make([]C.float, len(verts)*3)
	for i, v := range verts {
		flatVerts[i*3] = C.float(v[0])
		flatVerts[i*3+1] = C.float(v[1])
		flatVerts[i*3+2] = C.float(v[2])
	}
	flatIdx := make([]C.int, faceCount*3)
	for i := 0; i < faceCount; i++ {
		var idx [3]uint32
		binary.Read(r, binary.LittleEndian, &idx)
		r.Seek(skipBytes, io.SeekCurrent)
		flatIdx[i*3] = C.int(idx[0])
		flatIdx[i*3+1] = C.int(idx[1])
		flatIdx[i*3+2] = C.int(idx[2])
	}
	nv := C.int(len(verts))
	nt := C.int(faceCount)
	C.cm_shape_add_trimesh(s, &flatVerts[0], nv, &flatIdx[0], nt, 0)
	C.cm_shape_add_trimesh(s, &flatVerts[0], nv, &flatIdx[0], nt, 1)
	return true
}
