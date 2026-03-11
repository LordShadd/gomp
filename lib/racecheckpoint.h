#ifndef RACECHECKPOINT_H
#define RACECHECKPOINT_H

#include "ompcapi.h"
#include <stdbool.h>
#include <stdint.h>

bool RaceCheckpoint_Set(void *player, int type, float x, float y, float z, float nextX, float nextY, float nextZ, float radius);
bool RaceCheckpoint_Disable(void *player);
bool RaceCheckpoint_IsPlayerIn(void *player);
bool RaceCheckpoint_IsActive(void *player);
bool RaceCheckpoint_Get(void *player, float *x, float *y, float *z, float *nextX, float *nextY, float *nextZ, float *radius);

#endif
