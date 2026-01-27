#include "player.h"
#include "main.h"

int Player_NetStatsBytesReceived(void *player) {
  return api.Player.NetStatsBytesReceived(player);
}

int Player_NetStatsBytesSent(void *player) {
  return api.Player.NetStatsBytesSent(player);
}

int Player_NetStatsConnectionStatus(void *player) {
  return api.Player.NetStatsConnectionStatus(player);
}

int Player_NetStatsGetConnectedTime(void *player) {
  return api.Player.NetStatsGetConnectedTime(player);
}

bool Player_NetStatsGetIpPort(void *player, struct CAPIStringBuffer *output) {
  return api.Player.NetStatsGetIpPort(player, output);
}

int Player_NetStatsMessagesReceived(void *player) {
  return api.Player.NetStatsMessagesReceived(player);
}

int Player_NetStatsMessagesRecvPerSecond(void *player) {
  return api.Player.NetStatsMessagesRecvPerSecond(player);
}

int Player_NetStatsMessagesSent(void *player) {
  return api.Player.NetStatsMessagesSent(player);
}

float Player_NetStatsPacketLossPercent(void *player) {
  return api.Player.NetStatsPacketLossPercent(player);
}

int Player_GetCustomSkin(void *player) {
  return api.Player.GetCustomSkin(player);
}

int Player_GetDialog(void *player) { return api.Player.GetDialog(player); }

bool Player_GetDialogData(void *player, int *dialogid, int *style,
                          struct CAPIStringView *title,
                          struct CAPIStringView *body,
                          struct CAPIStringView *button1,
                          struct CAPIStringView *button2) {
  return api.Player.GetDialogData(player, dialogid, style, title, body, button1,
                                  button2);
}

void *Player_GetMenu(void *player) { return api.Player.GetMenu(player); }

void *Player_GetSurfingPlayerObject(void *player) {
  return api.Player.GetSurfingPlayerObject(player);
}

void *Player_GetCameraTargetPlayerObject(void *player) {
  return api.Player.GetCameraTargetPlayerObject(player);
}

void *Player_FromID(int playerid) { return api.Player.FromID(playerid); }

int Player_GetID(void *player) { return api.Player.GetID(player); }

bool Player_SendClientMessage(void *player, uint32_t color, const char *text) {
  return api.Player.SendClientMessage(player, color, text);
}

bool Player_SetCameraPos(void *player, float x, float y, float z) {
  return api.Player.SetCameraPos(player, x, y, z);
}

bool Player_SetDrunkLevel(void *player, int level) {
  return api.Player.SetDrunkLevel(player, level);
}

bool Player_SetInterior(void *player, int interior) {
  return api.Player.SetInterior(player, interior);
}

bool Player_SetWantedLevel(void *player, int level) {
  return api.Player.SetWantedLevel(player, level);
}

bool Player_SetWeather(void *player, int weather) {
  return api.Player.SetWeather(player, weather);
}

int Player_GetWeather(void *player) { return api.Player.GetWeather(player); }

bool Player_SetSkin(void *player, int skin) {
  return api.Player.SetSkin(player, skin);
}

bool Player_SetShopName(void *player, const char *name) {
  return api.Player.SetShopName(player, name);
}

bool Player_GiveMoney(void *player, int amount) {
  return api.Player.GiveMoney(player, amount);
}

bool Player_SetCameraLookAt(void *player, float x, float y, float z,
                            int cutType) {
  return api.Player.SetCameraLookAt(player, x, y, z, cutType);
}

bool Player_SetCameraBehind(void *player) {
  return api.Player.SetCameraBehind(player);
}

bool Player_CreateExplosion(void *player, float x, float y, float z, int type,
                            float radius) {
  return api.Player.CreateExplosion(player, x, y, z, type, radius);
}

bool Player_PlayAudioStream(void *player, const char *url, float x, float y,
                            float z, float distance, bool usePos) {
  return api.Player.PlayAudioStream(player, url, x, y, z, distance, usePos);
}

bool Player_StopAudioStream(void *player) {
  return api.Player.StopAudioStream(player);
}

bool Player_ToggleWidescreen(void *player, bool enable) {
  return api.Player.ToggleWidescreen(player, enable);
}

bool Player_IsWidescreenToggled(void *player) {
  return api.Player.IsWidescreenToggled(player);
}

bool Player_SetHealth(void *player, float health) {
  return api.Player.SetHealth(player, health);
}

float Player_GetHealth(void *player) { return api.Player.GetHealth(player); }

bool Player_SetArmor(void *player, float armor) {
  return api.Player.SetArmor(player, armor);
}

float Player_GetArmor(void *player) { return api.Player.GetArmor(player); }

bool Player_SetTeam(void *player, int team) {
  return api.Player.SetTeam(player, team);
}

int Player_GetTeam(void *player) { return api.Player.GetTeam(player); }

bool Player_SetScore(void *player, int score) {
  return api.Player.SetScore(player, score);
}

int Player_GetScore(void *player) { return api.Player.GetScore(player); }

int Player_GetSkin(void *player) { return api.Player.GetSkin(player); }

bool Player_SetColor(void *player, uint32_t color) {
  return api.Player.SetColor(player, color);
}

uint32_t Player_GetColor(void *player) { return api.Player.GetColor(player); }

uint32_t Player_GetDefaultColor(void *player) {
  return api.Player.GetDefaultColor(player);
}

int Player_GetDrunkLevel(void *player) {
  return api.Player.GetDrunkLevel(player);
}

bool Player_GiveWeapon(void *player, int weapon, int ammo) {
  return api.Player.GiveWeapon(player, weapon, ammo);
}

bool Player_RemoveWeapon(void *player, int weapon) {
  return api.Player.RemoveWeapon(player, weapon);
}

int Player_GetMoney(void *player) { return api.Player.GetMoney(player); }

bool Player_ResetMoney(void *player) { return api.Player.ResetMoney(player); }

int Player_SetName(void *player, const char *name) {
  return api.Player.SetName(player, name);
}

int Player_GetName(void *player, struct CAPIStringView *name) {
  return api.Player.GetName(player, name);
}

int Player_GetState(void *player) { return api.Player.GetState(player); }

int Player_GetPing(void *player) { return api.Player.GetPing(player); }

int Player_GetWeapon(void *player) { return api.Player.GetWeapon(player); }

bool Player_SetTime(void *player, int hour, int minute) {
  return api.Player.SetTime(player, hour, minute);
}

bool Player_GetTime(void *player, int *hour, int *minute) {
  return api.Player.GetTime(player, hour, minute);
}

bool Player_ToggleClock(void *player, bool enable) {
  return api.Player.ToggleClock(player, enable);
}

bool Player_HasClock(void *player) { return api.Player.HasClock(player); }

bool Player_ForceClassSelection(void *player) {
  return api.Player.ForceClassSelection(player);
}

int Player_GetWantedLevel(void *player) {
  return api.Player.GetWantedLevel(player);
}

bool Player_SetFightingStyle(void *player, int style) {
  return api.Player.SetFightingStyle(player, style);
}

int Player_GetFightingStyle(void *player) {
  return api.Player.GetFightingStyle(player);
}

bool Player_SetVelocity(void *player, float x, float y, float z) {
  return api.Player.SetVelocity(player, x, y, z);
}

bool Player_GetVelocity(void *player, float *x, float *y, float *z) {
  return api.Player.GetVelocity(player, x, y, z);
}

bool Player_GetCameraPos(void *player, float *x, float *y, float *z) {
  return api.Player.GetCameraPos(player, x, y, z);
}

float Player_GetDistanceFromPoint(void *player, float x, float y, float z) {
  return api.Player.GetDistanceFromPoint(player, x, y, z);
}

int Player_GetInterior(void *player) { return api.Player.GetInterior(player); }

bool Player_SetPos(void *player, float x, float y, float z) {
  return api.Player.SetPos(player, x, y, z);
}

bool Player_GetPos(void *player, float *x, float *y, float *z) {
  return api.Player.GetPos(player, x, y, z);
}

int Player_GetVirtualWorld(void *player) {
  return api.Player.GetVirtualWorld(player);
}

bool Player_IsNPC(void *player) { return api.Player.IsNPC(player); }

bool Player_IsStreamedIn(void *player, void *other) {
  return api.Player.IsStreamedIn(player, other);
}

bool Player_PlayGameSound(void *player, int sound, float x, float y, float z) {
  return api.Player.PlayGameSound(player, sound, x, y, z);
}

bool Player_SpectatePlayer(void *player, void *target, int mode) {
  return api.Player.SpectatePlayer(player, target, mode);
}

bool Player_SpectateVehicle(void *player, void *target, int mode) {
  return api.Player.SpectateVehicle(player, target, mode);
}

bool Player_SetVirtualWorld(void *player, int vw) {
  return api.Player.SetVirtualWorld(player, vw);
}

bool Player_SetWorldBounds(void *player, float xMax, float xMin, float yMax,
                           float yMin) {
  return api.Player.SetWorldBounds(player, xMax, xMin, yMax, yMin);
}

bool Player_ClearWorldBounds(void *player) {
  return api.Player.ClearWorldBounds(player);
}

bool Player_GetWorldBounds(void *player, float *xmax, float *xmin, float *ymax,
                           float *ymin) {
  return api.Player.GetWorldBounds(player, xmax, xmin, ymax, ymin);
}

bool Player_ClearAnimations(void *player, int syncType) {
  return api.Player.ClearAnimations(player, syncType);
}

bool Player_GetLastShotVectors(void *player, float *origin_x, float *origin_y,
                               float *origin_z, float *hit_x, float *hit_y,
                               float *hit_z) {
  return api.Player.GetLastShotVectors(player, origin_x, origin_y, origin_z,
                                       hit_x, hit_y, hit_z);
}

void *Player_GetCameraTargetPlayer(void *player) {
  return api.Player.GetCameraTargetPlayer(player);
}

void *Player_GetCameraTargetActor(void *player) {
  return api.Player.GetCameraTargetActor(player);
}

void *Player_GetCameraTargetObject(void *player) {
  return api.Player.GetCameraTargetObject(player);
}

void *Player_GetCameraTargetVehicle(void *player) {
  return api.Player.GetCameraTargetVehicle(player);
}

bool Player_PutInVehicle(void *player, void *vehicle, int seat) {
  return api.Player.PutInVehicle(player, vehicle, seat);
}

bool Player_RemoveBuilding(void *player, int model, float x, float y, float z,
                           float radius) {
  return api.Player.RemoveBuilding(player, model, x, y, z, radius);
}

int Player_GetBuildingsRemoved(void *player) {
  return api.Player.GetBuildingsRemoved(player);
}

bool Player_RemoveFromVehicle(void *player, bool force) {
  return api.Player.RemoveFromVehicle(player, force);
}

bool Player_RemoveMapIcon(void *player, int icon) {
  return api.Player.RemoveMapIcon(player, icon);
}

bool Player_SetMapIcon(void *player, int iconID, float x, float y, float z,
                       int type, uint32_t color, int style) {
  return api.Player.SetMapIcon(player, iconID, x, y, z, type, color, style);
}

bool Player_ResetWeapons(void *player) {
  return api.Player.ResetWeapons(player);
}

bool Player_SetAmmo(void *player, uint8_t id, uint32_t ammo) {
  return api.Player.SetAmmo(player, id, ammo);
}

bool Player_SetArmedWeapon(void *player, uint8_t weapon) {
  return api.Player.SetArmedWeapon(player, weapon);
}

bool Player_SetChatBubble(void *player, const char *text, uint32_t color,
                          float drawdistance, int expiretime) {
  return api.Player.SetChatBubble(player, text, color, drawdistance,
                                  expiretime);
}

bool Player_SetPosFindZ(void *player, float x, float y, float z) {
  return api.Player.SetPosFindZ(player, x, y, z);
}

bool Player_SetSkillLevel(void *player, uint8_t weapon, int level) {
  return api.Player.SetSkillLevel(player, weapon, level);
}

bool Player_SetSpecialAction(void *player, uint32_t action) {
  return api.Player.SetSpecialAction(player, action);
}

bool Player_ShowNameTagForPlayer(void *player, void *other, bool enable) {
  return api.Player.ShowNameTagForPlayer(player, other, enable);
}

bool Player_ToggleControllable(void *player, bool enable) {
  return api.Player.ToggleControllable(player, enable);
}

bool Player_ToggleSpectating(void *player, bool enable) {
  return api.Player.ToggleSpectating(player, enable);
}

bool Player_ApplyAnimation(void *player, const char *animlib,
                           const char *animname, float delta, bool loop,
                           bool lockX, bool lockY, bool freeze, uint32_t time,
                           int sync) {
  return api.Player.ApplyAnimation(player, animlib, animname, delta, loop,
                                   lockX, lockY, freeze, time, sync);
}

bool Player_GetAnimationName(int index, struct CAPIStringView *lib,
                             struct CAPIStringView *name) {
  return api.Player.GetAnimationName(index, lib, name);
}

bool Player_EditAttachedObject(void *player, int index) {
  return api.Player.EditAttachedObject(player, index);
}

bool Player_EnableCameraTarget(void *player, bool enable) {
  return api.Player.EnableCameraTarget(player, enable);
}

bool Player_EnableStuntBonus(void *player, bool enable) {
  return api.Player.EnableStuntBonus(player, enable);
}

int Player_GetPlayerAmmo(void *player) {
  return api.Player.GetPlayerAmmo(player);
}

int Player_GetAnimationIndex(void *player) {
  return api.Player.GetAnimationIndex(player);
}

float Player_GetFacingAngle(void *player) {
  return api.Player.GetFacingAngle(player);
}

int Player_GetIp(void *player, struct CAPIStringBuffer *ip) {
  return api.Player.GetIp(player, ip);
}

int Player_GetSpecialAction(void *player) {
  return api.Player.GetSpecialAction(player);
}

int Player_GetVehicleID(void *player) {
  return api.Player.GetVehicleID(player);
}

int Player_GetVehicleSeat(void *player) {
  return api.Player.GetVehicleSeat(player);
}

bool Player_GetWeaponData(void *player, int slot, int *weaponid, int *ammo) {
  return api.Player.GetWeaponData(player, slot, weaponid, ammo);
}

int Player_GetWeaponState(void *player) {
  return api.Player.GetWeaponState(player);
}

bool Player_InterpolateCameraPos(void *player, float from_x, float from_y,
                                 float from_z, float to_x, float to_y,
                                 float to_z, int time, int cut) {
  return api.Player.InterpolateCameraPos(player, from_x, from_y, from_z, to_x,
                                         to_y, to_z, time, cut);
}

bool Player_InterpolateCameraLookAt(void *player, float from_x, float from_y,
                                    float from_z, float to_x, float to_y,
                                    float to_z, int time, int cut) {
  return api.Player.InterpolateCameraLookAt(player, from_x, from_y, from_z,
                                            to_x, to_y, to_z, time, cut);
}

bool Player_IsPlayerAttachedObjectSlotUsed(void *player, int index) {
  return api.Player.IsPlayerAttachedObjectSlotUsed(player, index);
}

bool Player_AttachCameraToObject(void *player, void *object) {
  return api.Player.AttachCameraToObject(player, object);
}

bool Player_AttachCameraToPlayerObject(void *player, void *object) {
  return api.Player.AttachCameraToPlayerObject(player, object);
}

float Player_GetCameraAspectRatio(void *player) {
  return api.Player.GetCameraAspectRatio(player);
}

bool Player_GetCameraFrontVector(void *player, float *x, float *y, float *z) {
  return api.Player.GetCameraFrontVector(player, x, y, z);
}

int Player_GetCameraMode(void *player) {
  return api.Player.GetCameraMode(player);
}

bool Player_GetKeys(void *player, int *keys, int *updown, int *leftright) {
  return api.Player.GetKeys(player, keys, updown, leftright);
}

void *Player_GetSurfingVehicle(void *player) {
  return api.Player.GetSurfingVehicle(player);
}

void *Player_GetSurfingObject(void *player) {
  return api.Player.GetSurfingObject(player);
}

void *Player_GetTargetPlayer(void *player) {
  return api.Player.GetTargetPlayer(player);
}

void *Player_GetTargetActor(void *player) {
  return api.Player.GetTargetActor(player);
}

bool Player_IsInVehicle(void *player, void *targetVehicle) {
  return api.Player.IsInVehicle(player, targetVehicle);
}

bool Player_IsInAnyVehicle(void *player) {
  return api.Player.IsInAnyVehicle(player);
}

bool Player_IsInRangeOfPoint(void *player, float range, float x, float y,
                             float z) {
  return api.Player.IsInRangeOfPoint(player, range, x, y, z);
}

bool Player_PlayCrimeReport(void *player, void *suspect, int crime) {
  return api.Player.PlayCrimeReport(player, suspect, crime);
}

bool Player_RemoveAttachedObject(void *player, int index) {
  return api.Player.RemoveAttachedObject(player, index);
}

bool Player_SetAttachedObject(void *player, int index, int modelid, int bone,
                              float offsetX, float offsetY, float offsetZ,
                              float rotationX, float rotationY, float rotationZ,
                              float scaleX, float scaleY, float scaleZ,
                              uint32_t materialcolor1,
                              uint32_t materialcolor2) {
  return api.Player.SetAttachedObject(player, index, modelid, bone, offsetX,
                                      offsetY, offsetZ, rotationX, rotationY,
                                      rotationZ, scaleX, scaleY, scaleZ,
                                      materialcolor1, materialcolor2);
}

bool Player_GetAttachedObject(void *player, int index, int *modelid, int *bone,
                              float *offsetX, float *offsetY, float *offsetZ,
                              float *rotationX, float *rotationY,
                              float *rotationZ, float *scaleX, float *scaleY,
                              float *scaleZ, int *materialcolor1,
                              int *materialcolor2) {
  return api.Player.GetAttachedObject(player, index, modelid, bone, offsetX,
                                      offsetY, offsetZ, rotationX, rotationY,
                                      rotationZ, scaleX, scaleY, scaleZ,
                                      materialcolor1, materialcolor2);
}

bool Player_SetFacingAngle(void *player, float angle) {
  return api.Player.SetFacingAngle(player, angle);
}

bool Player_SetMarkerForPlayer(void *player, void *other, uint32_t color) {
  return api.Player.SetMarkerForPlayer(player, other, color);
}

uint32_t Player_GetMarkerForPlayer(void *player, void *other) {
  return api.Player.GetMarkerForPlayer(player, other);
}

bool Player_AllowTeleport(void *player, bool allow) {
  return api.Player.AllowTeleport(player, allow);
}

bool Player_IsTeleportAllowed(void *player) {
  return api.Player.IsTeleportAllowed(player);
}

bool Player_DisableRemoteVehicleCollisions(void *player, bool disable) {
  return api.Player.DisableRemoteVehicleCollisions(player, disable);
}

float Player_GetCameraZoom(void *player) {
  return api.Player.GetCameraZoom(player);
}

bool Player_SelectTextDraw(void *player, uint32_t hoverColour) {
  return api.Player.SelectTextDraw(player, hoverColour);
}

bool Player_CancelSelectTextDraw(void *player) {
  return api.Player.CancelSelectTextDraw(player);
}

bool Player_SendClientCheck(void *player, int actionType, int address,
                            int offset, int count) {
  return api.Player.SendClientCheck(player, actionType, address, offset, count);
}

bool Player_Spawn(void *player) { return api.Player.Spawn(player); }

bool Player_GPCI(void *player, struct CAPIStringView *gpci) {
  return api.Player.GPCI(player, gpci);
}

bool Player_IsAdmin(void *player) { return api.Player.IsAdmin(player); }

bool Player_Kick(void *player) { return api.Player.Kick(player); }

bool Player_ShowGameText(void *player, const char *text, int time, int style) {
  return api.Player.ShowGameText(player, text, time, style);
}

bool Player_HideGameText(void *player, int style) {
  return api.Player.HideGameText(player, style);
}

bool Player_HasGameText(void *player, int style) {
  return api.Player.HasGameText(player, style);
}

bool Player_GetGameText(void *player, int style, struct CAPIStringView *message,
                        int *time, int *remaining) {
  return api.Player.GetGameText(player, style, message, time, remaining);
}

bool Player_Ban(void *player) { return api.Player.Ban(player); }

bool Player_BanEx(void *player, const char *reason) {
  return api.Player.BanEx(player, reason);
}

bool Player_SendDeathMessage(void *player, void *killer, void *killee,
                             int weapon) {
  return api.Player.SendDeathMessage(player, killer, killee, weapon);
}

bool Player_SendMessageToPlayer(void *player, void *sender,
                                const char *message) {
  return api.Player.SendMessageToPlayer(player, sender, message);
}

int Player_GetVersion(void *player, struct CAPIStringView *version) {
  return api.Player.GetVersion(player, version);
}

int Player_GetSkillLevel(void *player, int skill) {
  return api.Player.GetSkillLevel(player, skill);
}

float Player_GetZAim(void *player) { return api.Player.GetZAim(player); }

bool Player_GetSurfingOffsets(void *player, float *offsetX, float *offsetY,
                              float *offsetZ) {
  return api.Player.GetSurfingOffsets(player, offsetX, offsetY, offsetZ);
}

bool Player_GetRotationQuat(void *player, float *x, float *y, float *z,
                            float *w) {
  return api.Player.GetRotationQuat(player, x, y, z, w);
}

int Player_GetPlayerSpectateID(void *player) {
  return api.Player.GetPlayerSpectateID(player);
}

int Player_GetSpectateType(void *player) {
  return api.Player.GetSpectateType(player);
}

uint32_t Player_GetRawIp(void *player) { return api.Player.GetRawIp(player); }

bool Player_SetGravity(void *player, float gravity) {
  return api.Player.SetGravity(player, gravity);
}

float Player_GetGravity(void *player) { return api.Player.GetGravity(player); }

bool Player_SetAdmin(void *player, bool set) {
  return api.Player.SetAdmin(player, set);
}

bool Player_IsSpawned(void *player) { return api.Player.IsSpawned(player); }

bool Player_IsControllable(void *player) {
  return api.Player.IsControllable(player);
}

bool Player_IsCameraTargetEnabled(void *player) {
  return api.Player.IsCameraTargetEnabled(player);
}

bool Player_ToggleGhostMode(void *player, bool toggle) {
  return api.Player.ToggleGhostMode(player, toggle);
}

bool Player_GetGhostMode(void *player) {
  return api.Player.GetGhostMode(player);
}

bool Player_AllowWeapons(void *player, bool allow) {
  return api.Player.AllowWeapons(player, allow);
}

bool Player_AreWeaponsAllowed(void *player) {
  return api.Player.AreWeaponsAllowed(player);
}

bool Player_IsPlayerUsingOfficialClient(void *player) {
  return api.Player.IsPlayerUsingOfficialClient(player);
}

int Player_GetAnimationFlags(void *player) {
  return api.Player.GetAnimationFlags(player);
}

bool Player_IsInDriveByMode(void *player) {
  return api.Player.IsInDriveByMode(player);
}

bool Player_IsCuffed(void *player) { return api.Player.IsCuffed(player); }

bool Player_IsUsingOmp(void *player) { return api.Player.IsUsingOmp(player); }

bool Player_IsInModShop(void *player) { return api.Player.IsInModShop(player); }

int Player_GetSirenState(void *player) {
  return api.Player.GetSirenState(player);
}

int Player_GetLandingGearState(void *player) {
  return api.Player.GetLandingGearState(player);
}

uint32_t Player_GetHydraReactorAngle(void *player) {
  return api.Player.GetHydraReactorAngle(player);
}

float Player_GetTrainSpeed(void *player) {
  return api.Player.GetTrainSpeed(player);
