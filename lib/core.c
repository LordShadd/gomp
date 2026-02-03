#include "core.h"
#include "main.h"

uint32_t Core_TickCount() { return api.Core.TickCount(); }

int Core_MaxPlayers() { return api.Core.MaxPlayers(); }

bool Core_Log(const char *text) { return api.Core.Log(text); }

bool Core_IsAdminTeleportAllowed() { return api.Core.IsAdminTeleportAllowed(); }

bool Core_AllowAdminTeleport(bool allow) {
  return api.Core.AllowAdminTeleport(allow);
}

bool Core_AreAllAnimationsEnabled() {
  return api.Core.AreAllAnimationsEnabled();
}

bool Core_EnableAllAnimations(bool allow) {
  return api.Core.EnableAllAnimations(allow);
}

bool Core_IsAnimationLibraryValid(const char *name) {
  return api.Core.IsAnimationLibraryValid(name);
}

bool Core_AreInteriorWeaponsAllowed() {
  return api.Core.AreInteriorWeaponsAllowed();
}

bool Core_AllowInteriorWeapons(bool allow) {
  return api.Core.AllowInteriorWeapons(allow);
}

bool Core_BlockIpAddress(const char *ipAddress, int timeMS) {
  return api.Core.BlockIpAddress(ipAddress, timeMS);
}

bool Core_UnBlockIpAddress(const char *ipAddress) {
  return api.Core.UnBlockIpAddress(ipAddress);
}

bool Core_DisableEntryExitMarkers() {
  return api.Core.DisableEntryExitMarkers();
}

bool Core_DisableNameTagsLOS() { return api.Core.DisableNameTagsLOS(); }

bool Core_EnableZoneNames(bool enable) {
  return api.Core.EnableZoneNames(enable);
}

bool Core_ShowGameTextForAll(const char *msg, int time, int style) {
  return api.Core.ShowGameTextForAll(msg, time, style);
}

bool Core_HideGameTextForAll(int style) {
  return api.Core.HideGameTextForAll(style);
}

int Core_NetworkStats(struct CAPIStringBuffer *output) {
  return api.Core.NetworkStats(output);
}

int Core_ServerTickRate() { return api.Core.ServerTickRate(); }

bool Core_GetWeaponName(int weaponid, struct CAPIStringView *output) {
  return api.Core.GetWeaponName(weaponid, output);
}

bool Core_SetChatRadius(float globalChatRadius) {
  return api.Core.SetChatRadius(globalChatRadius);
}

bool Core_SetMarkerRadius(float playerMarkerRadius) {
  return api.Core.SetMarkerRadius(playerMarkerRadius);
}

bool Core_SendRconCommand(const char *command) {
  return api.Core.SendRconCommand(command);
}

bool Core_SetDeathDropAmount(int amount) {
  return api.Core.SetDeathDropAmount(amount);
}

bool Core_GameMode_SetText(const char *string) {
  return api.Core.GameMode_SetText(string);
}

bool Core_SetGravity(float gravity) { return api.Core.SetGravity(gravity); }

float Core_GetGravity() { return api.Core.GetGravity(); }

bool Core_SetNameTagsDrawDistance(float distance) {
  return api.Core.SetNameTagsDrawDistance(distance);
}

bool Core_SetWeather(int weatherid) { return api.Core.SetWeather(weatherid); }

bool Core_SetWorldTime(int hour) { return api.Core.SetWorldTime(hour); }

bool Core_ShowNameTags(bool show) { return api.Core.ShowNameTags(show); }

bool Core_ShowPlayerMarkers(int mode) {
  return api.Core.ShowPlayerMarkers(mode);
}

bool Core_UsePedAnims() { return api.Core.UsePedAnims(); }

int Core_GetWeather() { return api.Core.GetWeather(); }

int Core_GetWorldTime() { return api.Core.GetWorldTime(); }

bool Core_ToggleChatTextReplacement(bool enable) {
  return api.Core.ToggleChatTextReplacement(enable);
}

bool Core_IsChatTextReplacementToggled() {
  return api.Core.IsChatTextReplacementToggled();
}

bool Core_IsNickNameValid(const char *name) {
  return api.Core.IsNickNameValid(name);
}

bool Core_AllowNickNameCharacter(int character, bool allow) {
  return api.Core.AllowNickNameCharacter(character, allow);
}

bool Core_IsNickNameCharacterAllowed(int character) {
  return api.Core.IsNickNameCharacterAllowed(character);
}

bool Core_ClearBanList() { return api.Core.ClearBanList(); }

bool Core_IsIpAddressBanned(const char *ip) {
  return api.Core.IsIpAddressBanned(ip);
}

int Core_GetWeaponSlot(uint8_t weapon) {
  return api.Core.GetWeaponSlot(weapon);
}

bool Core_AddRule(const char *name, const char *value) {
  return api.Core.AddRule(name, value);
}

bool Core_IsValidRule(const char *name) { return api.Core.IsValidRule(name); }

bool Core_RemoveRule(const char *name) { return api.Core.RemoveRule(name); }
