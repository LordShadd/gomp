#ifndef BULLET_CAPI_H
#define BULLET_CAPI_H

#ifdef __cplusplus
extern "C" {
#endif

typedef void* CMWorld;
typedef void* CMShape;

CMWorld cm_world_create();
void    cm_world_destroy(CMWorld w);

CMShape cm_shape_create();
void    cm_shape_destroy(CMShape s);
void    cm_shape_add_sphere(CMShape s, float offX, float offY, float offZ, float radius);
void    cm_shape_add_box(CMShape s, float cx, float cy, float cz, float sx, float sy, float sz);
void    cm_shape_add_trimesh(CMShape s, float *verts, int numVerts, int *indices, int numTris, int convex);
void    cm_shape_bounding_sphere(CMShape s, float *cx, float *cy, float *cz, float *r);
void    cm_shape_bounding_box(CMShape s, float *minx, float *miny, float *minz,
                               float *maxx, float *maxy, float *maxz);

int  cm_world_add_body(CMWorld w, CMShape s,
                        float x, float y, float z,
                        float qx, float qy, float qz, float qw,
                        int modelId, int bodyTag);
void cm_world_remove_body(CMWorld w, int handle);
int  cm_world_set_body_pos(CMWorld w, int handle, float x, float y, float z);
int  cm_world_set_body_rot(CMWorld w, int handle, float qx, float qy, float qz, float qw);

int  cm_raytest(CMWorld w,
                float sx, float sy, float sz,
                float ex, float ey, float ez,
                float *rx, float *ry, float *rz,
                float *nx, float *ny, float *nz,
                int *modelId, int *bodyTag);

int  cm_raytest_ex(CMWorld w,
                   float sx, float sy, float sz,
                   float ex, float ey, float ez,
                   float *rx, float *ry, float *rz,
                   float *qx, float *qy, float *qz, float *qw,
                   float *px, float *py, float *pz,
                   int *modelId);

int  cm_raytest_all(CMWorld w,
                    float sx, float sy, float sz,
                    float ex, float ey, float ez,
                    float *rxArr, float *ryArr, float *rzArr,
                    float *distArr, int *modelArr, int *bodyTagArr,
                    int maxSize);

int  cm_contact_test(CMWorld w, CMShape s,
                      float x, float y, float z,
                      float qx, float qy, float qz, float qw);

CMShape cm_load_dff(const char *dffPath, int modelid);

#ifdef __cplusplus
}
#endif

#endif
