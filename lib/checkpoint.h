#ifndef CHECKPOINT_H
#define CHECKPOINT_H

#include "ompcapi.h"
#include <stdbool.h>
#include <stdint.h>

bool Checkpoint_Set(void *player, float x, float y, float z, float radius);
bool Checkpoint_Disable(void *player);
bool Checkpoint_IsPlayerIn(void *player);
bool Checkpoint_IsActive(void *player);
bool Checkpoint_Get(void *player, float *x, float *y, float *z, float *radius);

#endif
