#include "checkpoint.h"
#include "main.h"

bool Checkpoint_Set(void *player, float x, float y, float z, float radius) { return api.Checkpoint.Set(player, x, y, z, radius); }
bool Checkpoint_Disable(void *player) { return api.Checkpoint.Disable(player); }
bool Checkpoint_IsPlayerIn(void *player) { return api.Checkpoint.IsPlayerIn(player); }
bool Checkpoint_IsActive(void *player) { return api.Checkpoint.IsActive(player); }
bool Checkpoint_Get(void *player, float *x, float *y, float *z, float *radius) { return api.Checkpoint.Get(player, x, y, z, radius); }
