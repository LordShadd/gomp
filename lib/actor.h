#ifndef ACTOR_H
#define ACTOR_H

#include "ompcapi.h"
#include <stdbool.h>
#include <stdint.h>

void *Actor_Create(int model, float x, float y, float z, float rot, int *id);

bool Actor_Destroy(void *actor);

void *Actor_FromID(int actorid);

int Actor_GetID(void *actor);

bool Actor_IsStreamedInFor(void *actor, void *player);

bool Actor_SetVirtualWorld(void *actor, int vw);

int Actor_GetVirtualWorld(void *actor);

bool Actor_ApplyAnimation(void *actor, const char *name, const char *library,
                          float delta, bool loop, bool lockX, bool lockY,
                          bool freeze, int time);

bool Actor_ClearAnimations(void *actor);

bool Actor_SetPos(void *actor, float x, float y, float z);

bool Actor_GetPos(void *actor, float *x, float *y, float *z);

bool Actor_SetFacingAngle(void *actor, float angle);

float Actor_GetFacingAngle(void *actor);

bool Actor_SetHealth(void *actor, float hp);

float Actor_GetHealth(void *actor);

bool Actor_SetInvulnerable(void *actor, bool toggle);

bool Actor_IsInvulnerable(void *actor);

bool Actor_IsValid(void *actor);

bool Actor_SetSkin(void *actor, int skin);

int Actor_GetSkin(void *actor);

bool Actor_GetAnimation(void *actor, struct CAPIStringView *library,
                        struct CAPIStringView *name, float *delta, bool *loop,
                        bool *lockX, bool *lockY, bool *freeze, int *time);

bool Actor_GetSpawnInfo(void *actor, float *x, float *y, float *z, float *angle,
                        int *skin);

#endif
