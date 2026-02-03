#ifndef CORE_H
#define CORE_H

#include "ompcapi.h"
#include <stdbool.h>
#include <stdint.h>

uint32_t Core_TickCount();
int Core_MaxPlayers();
bool Core_Log(const char *text);
bool Core_IsAdminTeleportAllowed();
bool Core_AllowAdminTeleport(bool allow);
bool Core_AreAllAnimationsEnabled();
bool Core_EnableAllAnimations(bool allow);
bool Core_IsAnimationLibraryValid(const char *name);
bool Core_AreInteriorWeaponsAllowed();
bool Core_AllowInteriorWeapons(bool allow);
bool Core_BlockIpAddress(const char *ipAddress, int timeMS);
bool Core_UnBlockIpAddress(const char *ipAddress);
bool Core_DisableEntryExitMarkers();
bool Core_DisableNameTagsLOS();
bool Core_EnableZoneNames(bool enable);
bool Core_ShowGameTextForAll(const char *msg, int time, int style);
bool Core_HideGameTextForAll(int style);
int Core_NetworkStats(struct CAPIStringBuffer *output);
int Core_ServerTickRate();
bool Core_GetWeaponName(int weaponid, struct CAPIStringView *output);
bool Core_SetChatRadius(float globalChatRadius);
bool Core_SetMarkerRadius(float playerMarkerRadius);
bool Core_SendRconCommand(const char *command);
bool Core_SetDeathDropAmount(int amount);
bool Core_GameMode_SetText(const char *string);
bool Core_SetGravity(float gravity);
float Core_GetGravity();
bool Core_SetNameTagsDrawDistance(float distance);
bool Core_SetWeather(int weatherid);
bool Core_SetWorldTime(int hour);
bool Core_ShowNameTags(bool show);
bool Core_ShowPlayerMarkers(int mode);
bool Core_UsePedAnims();
int Core_GetWeather();
int Core_GetWorldTime();
bool Core_ToggleChatTextReplacement(bool enable);
bool Core_IsChatTextReplacementToggled();
bool Core_IsNickNameValid(const char *name);
bool Core_AllowNickNameCharacter(int character, bool allow);
bool Core_IsNickNameCharacterAllowed(int character);
bool Core_ClearBanList();
bool Core_IsIpAddressBanned(const char *ip);
int Core_GetWeaponSlot(uint8_t weapon);
bool Core_AddRule(const char *name, const char *value);
bool Core_IsValidRule(const char *name);
bool Core_RemoveRule(const char *name);

#endif
