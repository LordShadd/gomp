#include "all.h"
#include "main.h"

bool All_SendClientMessage(uint32_t color, const char *text) {
  return api.All.SendClientMessage(color, text);
}

bool All_CreateExplosion(float x, float y, float z, int type, float radius) {
  return api.All.CreateExplosion(x, y, z, type, radius);
}

bool All_SendDeathMessage(void *killer, void *killee, int weapon) {
  return api.All.SendDeathMessage(killer, killee, weapon);
}

bool All_EnableStuntBonus(bool enable) {
  return api.All.EnableStuntBonus(enable);
}
