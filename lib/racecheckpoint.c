#include "racecheckpoint.h"
#include "main.h"

bool RaceCheckpoint_Set(void *player, int type, float x, float y, float z, float nextX, float nextY, float nextZ, float radius) { return api.RaceCheckpoint.Set(player, type, x, y, z, nextX, nextY, nextZ, radius); }
bool RaceCheckpoint_Disable(void *player) { return api.RaceCheckpoint.Disable(player); }
bool RaceCheckpoint_IsPlayerIn(void *player) { return api.RaceCheckpoint.IsPlayerIn(player); }
bool RaceCheckpoint_IsActive(void *player) { return api.RaceCheckpoint.IsActive(player); }
bool RaceCheckpoint_Get(void *player, float *x, float *y, float *z, float *nextX, float *nextY, float *nextZ, float *radius) { return api.RaceCheckpoint.Get(player, x, y, z, nextX, nextY, nextZ, radius); }
