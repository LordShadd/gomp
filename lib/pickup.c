#include "pickup.h"
#include "main.h"

void *Pickup_Create(int model, int type, float x, float y, float z,
                    int virtualWorld, int *id) {
  return api.Pickup.Create(model, type, x, y, z, virtualWorld, id);
}

bool Pickup_AddStatic(int model, int type, float x, float y, float z,
                      int virtualWorld) {
  return api.Pickup.AddStatic(model, type, x, y, z, virtualWorld);
}

bool Pickup_Destroy(void *pickup) { return api.Pickup.Destroy(pickup); }

void *Pickup_FromID(int pickupid) { return api.Pickup.FromID(pickupid); }

int Pickup_GetID(void *pickup) { return api.Pickup.GetID(pickup); }

bool Pickup_IsValid(void *pickup) { return api.Pickup.IsValid(pickup); }

bool Pickup_IsStreamedIn(void *player, void *pickup) {
  return api.Pickup.IsStreamedIn(player, pickup);
}

bool Pickup_GetPos(void *pickup, float *x, float *y, float *z) {
  return api.Pickup.GetPos(pickup, x, y, z);
}

int Pickup_GetModel(void *pickup) { return api.Pickup.GetModel(pickup); }

int Pickup_GetType(void *pickup) { return api.Pickup.GetType(pickup); }

int Pickup_GetVirtualWorld(void *pickup) {
  return api.Pickup.GetVirtualWorld(pickup);
}

bool Pickup_SetPos(void *pickup, float x, float y, float z, bool update) {
  return api.Pickup.SetPos(pickup, x, y, z, update);
}

bool Pickup_SetModel(void *pickup, int model, bool update) {
  return api.Pickup.SetModel(pickup, model, update);
}

bool Pickup_SetType(void *pickup, int type, bool update) {
  return api.Pickup.SetType(pickup, type, update);
}

bool Pickup_SetVirtualWorld(void *pickup, int virtualworld) {
  return api.Pickup.SetVirtualWorld(pickup, virtualworld);
}

bool Pickup_ShowForPlayer(void *player, void *pickup) {
  return api.Pickup.ShowForPlayer(player, pickup);
}

bool Pickup_HideForPlayer(void *player, void *pickup) {
  return api.Pickup.HideForPlayer(player, pickup);
}

bool Pickup_IsHiddenForPlayer(void *player, void *pickup) {
  return api.Pickup.IsHiddenForPlayer(player, pickup);
}
