#ifndef ALL_H
#define ALL_H

#include <stdbool.h>
#include <stdint.h>

bool All_SendClientMessage(uint32_t color, const char *text);
bool All_CreateExplosion(float x, float y, float z, int type, float radius);
bool All_SendDeathMessage(void *killer, void *killee, int weapon);
bool All_EnableStuntBonus(bool enable);

#endif
