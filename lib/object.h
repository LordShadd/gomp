#ifndef PLAYER_H
#define PLAYER_H

#include "main.h"
#include "ompcapi.h"
#include <stdbool.h>
#include <stdint.h>

void *Object_Create(int modelid, float x, float y, float z, float rotationX,
                    float rotationY, float rotationZ, float drawDistance,
                    int *id);
bool Object_Destroy(void *object);
void *Object_FromID(int objectid);
int Object_GetID(void *object);
bool Object_AttachToVehicle(void *object, void *vehicle, float offsetX,
                            float offsetY, float offsetZ, float rotationX,
                            float rotationY, float rotationZ);
bool Object_AttachToObject(void *object, void *objAttachedTo, float offsetX,
                           float offsetY, float offsetZ, float rotationX,
                           float rotationY, float rotationZ, bool syncRotation);
bool Object_AttachToPlayer(void *object, void *player, float offsetX,
                           float offsetY, float offsetZ, float rotationX,
                           float rotationY, float rotationZ);
bool Object_SetPos(void *object, float x, float y, float z);
bool Object_GetPos(void *object, float *x, float *y, float *z);
bool Object_SetRot(void *object, float rotationX, float rotationY,
                   float rotationZ);
bool Object_GetRot(void *object, float *rotationX, float *rotationY,
                   float *rotationZ);
int Object_GetModel(void *object);
bool Object_SetNoCameraCollision(void *object);
bool Object_IsValid(void *object);
int Object_Move(void *object, float x, float y, float z, float speed,
                float rotationX, float rotationY, float rotationZ);
bool Object_Stop(void *object);
bool Object_IsMoving(void *object);
bool Object_BeginEditing(void *player, void *object);
bool Object_BeginSelecting(void *player);
bool Object_EndEditing(void *player);
bool Object_SetMaterial(void *object, int materialIndex, int modelId,
                        const char *textureLibrary, const char *textureName,
                        uint32_t materialColor);
bool Object_SetMaterialText(void *object, const char *text, int materialIndex,
                            int materialSize, const char *fontface,
                            int fontsize, bool bold, uint32_t fontColor,
                            uint32_t backgroundColor, int textalignment);
bool Object_SetDefaultCameraCollision(bool disable);
float Object_GetDrawDistance(void *object);
float Object_GetMoveSpeed(void *object);
bool Object_GetMovingTargetPos(void *object, float *targetX, float *targetY,
                               float *targetZ);
bool Object_GetMovingTargetRot(void *object, float *rotationX, float *rotationY,
                               float *rotationZ);
bool Object_GetAttachedData(void *object, int *parentVehicle, int *parentObject,
                            int *parentPlayer);
bool Object_GetAttachedOffset(void *object, float *offsetX, float *offsetY,
                              float *offsetZ, float *rotationX,
                              float *rotationY, float *rotationZ);
bool Object_GetSyncRotation(void *object);
bool Object_IsMaterialSlotUsed(void *object, int materialIndex);
bool Object_GetMaterial(void *object, int materialIndex, int *modelid,
                        struct CAPIStringView *textureLibrary,
                        struct CAPIStringView *textureName, int *materialColor);
bool Object_GetMaterialText(void *object, int materialIndex,
                            struct CAPIStringView *text, int *materialSize,
                            struct CAPIStringView *fontFace, int *fontSize,
                            bool *bold, int *fontColor, int *backgroundColor,
                            int *textAlignment);
bool Object_IsObjectNoCameraCollision(void *object);
uint8_t Object_GetType(void *player, int objectid);

#endif
