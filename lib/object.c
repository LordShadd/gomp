#include "object.h"
#include "ompcapi.h"

void *Object_Create(int modelid, float x, float y, float z, float rotationX,
                    float rotationY, float rotationZ, float drawDistance,
                    int *id) {
  return api.Object.Create(modelid, x, y, z, rotationX, rotationY, rotationZ,
                           drawDistance, id);
}

bool Object_Destroy(void *object) { return api.Object.Destroy(object); }

void *Object_FromID(int objectid) { return api.Object.FromID(objectid); }

int Object_GetID(void *object) { return api.Object.GetID(object); }

bool Object_AttachToVehicle(void *object, void *vehicle, float offsetX,
                            float offsetY, float offsetZ, float rotationX,
                            float rotationY, float rotationZ) {
  return api.Object.AttachToVehicle(object, vehicle, offsetX, offsetY, offsetZ,
                                    rotationX, rotationY, rotationZ);
}

bool Object_AttachToObject(void *object, void *objAttachedTo, float offsetX,
                           float offsetY, float offsetZ, float rotationX,
                           float rotationY, float rotationZ,
                           bool syncRotation) {
  return api.Object.AttachToObject(object, objAttachedTo, offsetX, offsetY,
                                   offsetZ, rotationX, rotationY, rotationZ,
                                   syncRotation);
}

bool Object_AttachToPlayer(void *object, void *player, float offsetX,
                           float offsetY, float offsetZ, float rotationX,
                           float rotationY, float rotationZ) {
  return api.Object.AttachToPlayer(object, player, offsetX, offsetY, offsetZ,
                                   rotationX, rotationY, rotationZ);
}

bool Object_SetPos(void *object, float x, float y, float z) {
  return api.Object.SetPos(object, x, y, z);
}

bool Object_GetPos(void *object, float *x, float *y, float *z) {
  return api.Object.GetPos(object, x, y, z);
}

bool Object_SetRot(void *object, float rotationX, float rotationY,
                   float rotationZ) {
  return api.Object.SetRot(object, rotationX, rotationY, rotationZ);
}

bool Object_GetRot(void *object, float *rotationX, float *rotationY,
                   float *rotationZ) {
  return api.Object.GetRot(object, rotationX, rotationY, rotationZ);
}

int Object_GetModel(void *object) { return api.Object.GetModel(object); }

bool Object_SetNoCameraCollision(void *object) {
  return api.Object.SetNoCameraCollision(object);
}

bool Object_IsValid(void *object) { return api.Object.IsValid(object); }

int Object_Move(void *object, float x, float y, float z, float speed,
                float rotationX, float rotationY, float rotationZ) {
  return api.Object.Move(object, x, y, z, speed, rotationX, rotationY,
                         rotationZ);
}

bool Object_Stop(void *object) { return api.Object.Stop(object); }

bool Object_IsMoving(void *object) { return api.Object.IsMoving(object); }

bool Object_BeginEditing(void *player, void *object) {
  return api.Object.BeginEditing(player, object);
}

bool Object_BeginSelecting(void *player) {
  return api.Object.BeginSelecting(player);
}

bool Object_EndEditing(void *player) { return api.Object.EndEditing(player); }

bool Object_SetMaterial(void *object, int materialIndex, int modelId,
                        const char *textureLibrary, const char *textureName,
                        uint32_t materialColor) {
  return api.Object.SetMaterial(object, materialIndex, modelId, textureLibrary,
                                textureName, materialColor);
}

bool Object_SetMaterialText(void *object, const char *text, int materialIndex,
                            int materialSize, const char *fontface,
                            int fontsize, bool bold, uint32_t fontColor,
                            uint32_t backgroundColor, int textalignment) {
  return api.Object.SetMaterialText(object, text, materialIndex, materialSize,
                                    fontface, fontsize, bold, fontColor,
                                    backgroundColor, textalignment);
}

bool Object_SetDefaultCameraCollision(bool disable) {
  return api.Object.SetDefaultCameraCollision(disable);
}

float Object_GetDrawDistance(void *object) {
  return api.Object.GetDrawDistance(object);
}

float Object_GetMoveSpeed(void *object) {
  return api.Object.GetMoveSpeed(object);
}

bool Object_GetMovingTargetPos(void *object, float *targetX, float *targetY,
                               float *targetZ) {
  return api.Object.GetMovingTargetPos(object, targetX, targetY, targetZ);
}

bool Object_GetMovingTargetRot(void *object, float *rotationX, float *rotationY,
                               float *rotationZ) {
  return api.Object.GetMovingTargetRot(object, rotationX, rotationY, rotationZ);
}

bool Object_GetAttachedData(void *object, int *parentVehicle, int *parentObject,
                            int *parentPlayer) {
  return api.Object.GetAttachedData(object, parentVehicle, parentObject,
                                    parentPlayer);
}

bool Object_GetAttachedOffset(void *object, float *offsetX, float *offsetY,
                              float *offsetZ, float *rotationX,
                              float *rotationY, float *rotationZ) {
  return api.Object.GetAttachedOffset(object, offsetX, offsetY, offsetZ,
                                      rotationX, rotationY, rotationZ);
}

bool Object_GetSyncRotation(void *object) {
  return api.Object.GetSyncRotation(object);
}

bool Object_IsMaterialSlotUsed(void *object, int materialIndex) {
  return api.Object.IsMaterialSlotUsed(object, materialIndex);
}

bool Object_GetMaterial(void *object, int materialIndex, int *modelid,
                        struct CAPIStringView *textureLibrary,
                        struct CAPIStringView *textureName,
                        int *materialColor) {
  return api.Object.GetMaterial(object, materialIndex, modelid, textureLibrary,
                                textureName, materialColor);
}

bool Object_GetMaterialText(void *object, int materialIndex,
                            struct CAPIStringView *text, int *materialSize,
                            struct CAPIStringView *fontFace, int *fontSize,
                            bool *bold, int *fontColor, int *backgroundColor,
                            int *textAlignment) {
  return api.Object.GetMaterialText(object, materialIndex, text, materialSize,
                                    fontFace, fontSize, bold, fontColor,
                                    backgroundColor, textAlignment);
}

bool Object_IsObjectNoCameraCollision(void *object) {
  return api.Object.IsObjectNoCameraCollision(object);
}

uint8_t Object_GetType(void *player, int objectid) {
  return api.Object.GetType(player, objectid);
}
