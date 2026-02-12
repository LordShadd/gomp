#ifndef ALL_H
#define ALL_H

#include <stdbool.h>
#include <stdint.h>

void *Pickup_Create(int model, int type, float x, float y, float z,
                    int virtualWorld, int *id);
bool Pickup_AddStatic(int model, int type, float x, float y, float z,
                      int virtualWorld);
bool Pickup_Destroy(void *pickup);
void *Pickup_FromID(int pickupid);
int Pickup_GetID(void *pickup);
bool Pickup_IsValid(void *pickup);
bool Pickup_IsStreamedIn(void *player, void *pickup);
bool Pickup_GetPos(void *pickup, float *x, float *y, float *z);
int Pickup_GetModel(void *pickup);
int Pickup_GetType(void *pickup);
int Pickup_GetVirtualWorld(void *pickup);
bool Pickup_SetPos(void *pickup, float x, float y, float z, bool update);
bool Pickup_SetModel(void *pickup, int model, bool update);
bool Pickup_SetType(void *pickup, int type, bool update);
bool Pickup_SetVirtualWorld(void *pickup, int virtualworld);
bool Pickup_ShowForPlayer(void *player, void *pickup);
bool Pickup_HideForPlayer(void *player, void *pickup);
bool Pickup_IsHiddenForPlayer(void *player, void *pickup);

#endif
