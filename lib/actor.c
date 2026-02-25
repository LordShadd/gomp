#include "actor.h"
#include "main.h"

void *Actor_Create(int model, float x, float y, float z, float rot, int *id) {
  return api.Actor.Create(model, x, y, z, rot, id);
}

bool Actor_Destroy(void *actor) { return api.Actor.Destroy(actor); }

void *Actor_FromID(int actorid) { return api.Actor.FromID(actorid); }

int Actor_GetID(void *actor) { return api.Actor.GetID(actor); }

bool Actor_IsStreamedInFor(void *actor, void *player) {
  return api.Actor.IsStreamedInFor(actor, player);
}

bool Actor_SetVirtualWorld(void *actor, int vw) {
  return api.Actor.SetVirtualWorld(actor, vw);
}

int Actor_GetVirtualWorld(void *actor) {
  return api.Actor.GetVirtualWorld(actor);
}

bool Actor_ApplyAnimation(void *actor, const char *name, const char *library,
                          float delta, bool loop, bool lockX, bool lockY,
                          bool freeze, int time) {
  return api.Actor.ApplyAnimation(actor, name, library, delta, loop, lockX,
                                  lockY, freeze, time);
}

bool Actor_ClearAnimations(void *actor) {
  return api.Actor.ClearAnimations(actor);
}

bool Actor_SetPos(void *actor, float x, float y, float z) {
  return api.Actor.SetPos(actor, x, y, z);
}

bool Actor_GetPos(void *actor, float *x, float *y, float *z) {
  return api.Actor.GetPos(actor, x, y, z);
}

bool Actor_SetFacingAngle(void *actor, float angle) {
  return api.Actor.SetFacingAngle(actor, angle);
}

float Actor_GetFacingAngle(void *actor) {
  return api.Actor.GetFacingAngle(actor);
}

bool Actor_SetHealth(void *actor, float hp) {
  return api.Actor.SetHealth(actor, hp);
}

float Actor_GetHealth(void *actor) { return api.Actor.GetHealth(actor); }

bool Actor_SetInvulnerable(void *actor, bool toggle) {
  return api.Actor.SetInvulnerable(actor, toggle);
}

bool Actor_IsInvulnerable(void *actor) {
  return api.Actor.IsInvulnerable(actor);
}

bool Actor_IsValid(void *actor) { return api.Actor.IsValid(actor); }

bool Actor_SetSkin(void *actor, int skin) {
  return api.Actor.SetSkin(actor, skin);
}

int Actor_GetSkin(void *actor) { return api.Actor.GetSkin(actor); }

bool Actor_GetAnimation(void *actor, struct CAPIStringView *library,
                        struct CAPIStringView *name, float *delta, bool *loop,
                        bool *lockX, bool *lockY, bool *freeze, int *time) {
  return api.Actor.GetAnimation(actor, library, name, delta, loop, lockX, lockY,
                                freeze, time);
}

bool Actor_GetSpawnInfo(void *actor, float *x, float *y, float *z, float *angle,
                        int *skin) {
  return api.Actor.GetSpawnInfo(actor, x, y, z, angle, skin);
}
