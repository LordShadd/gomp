#include "vehicle.h"
#include "main.h"

void *Vehicle_Create(int modelid, float x, float y, float z, float rotation,
                     int color1, int color2, int respawnDelay, bool addSiren,
                     int *id) {
  return api.Vehicle.Create(modelid, x, y, z, rotation, color1, color2,
                            respawnDelay, addSiren, id);
}

bool Vehicle_Destroy(void *vehicle) { return api.Vehicle.Destroy(vehicle); }

void *Vehicle_FromID(int vehicleid) { return api.Vehicle.FromID(vehicleid); }

int Vehicle_GetID(void *vehicle) { return api.Vehicle.GetID(vehicle); }

int Vehicle_GetMaxPassengerSeats(int modelid) {
  return api.Vehicle.GetMaxPassengerSeats(modelid);
}

bool Vehicle_IsStreamedIn(void *vehicle, void *player) {
  return api.Vehicle.IsStreamedIn(vehicle, player);
}

bool Vehicle_GetPos(void *vehicle, float *x, float *y, float *z) {
  return api.Vehicle.GetPos(vehicle, x, y, z);
}

bool Vehicle_SetPos(void *vehicle, float x, float y, float z) {
  return api.Vehicle.SetPos(vehicle, x, y, z);
}

float Vehicle_GetZAngle(void *vehicle) {
  return api.Vehicle.GetZAngle(vehicle);
}

bool Vehicle_GetRotationQuat(void *vehicle, float *w, float *x, float *y,
                             float *z) {
  return api.Vehicle.GetRotationQuat(vehicle, w, x, y, z);
}

float Vehicle_GetDistanceFromPoint(void *vehicle, float x, float y, float z) {
  return api.Vehicle.GetDistanceFromPoint(vehicle, x, y, z);
}

bool Vehicle_SetZAngle(void *vehicle, float angle) {
  return api.Vehicle.SetZAngle(vehicle, angle);
}

bool Vehicle_SetParamsForPlayer(void *vehicle, void *player, int objective,
                                int doors) {
  return api.Vehicle.SetParamsForPlayer(vehicle, player, objective, doors);
}

bool Vehicle_UseManualEngineAndLights() {
  return api.Vehicle.UseManualEngineAndLights();
}

bool Vehicle_SetParamsEx(void *vehicle, int engine, int lights, int alarm,
                         int doors, int bonnet, int boot, int objective) {
  return api.Vehicle.SetParamsEx(vehicle, engine, lights, alarm, doors, bonnet,
                                 boot, objective);
}

bool Vehicle_GetParamsEx(void *vehicle, int *engine, int *lights, int *alarm,
                         int *doors, int *bonnet, int *boot, int *objective) {
  return api.Vehicle.GetParamsEx(vehicle, engine, lights, alarm, doors, bonnet,
                                 boot, objective);
}

int Vehicle_GetParamsSirenState(void *vehicle) {
  return api.Vehicle.GetParamsSirenState(vehicle);
}

bool Vehicle_SetParamsCarDoors(void *vehicle, int frontLeft, int frontRight,
                               int rearLeft, int rearRight) {
  return api.Vehicle.SetParamsCarDoors(vehicle, frontLeft, frontRight, rearLeft,
                                       rearRight);
}

bool Vehicle_GetParamsCarDoors(void *vehicle, int *frontLeft, int *frontRight,
                               int *rearLeft, int *rearRight) {
  return api.Vehicle.GetParamsCarDoors(vehicle, frontLeft, frontRight, rearLeft,
                                       rearRight);
}

bool Vehicle_SetParamsCarWindows(void *vehicle, int frontLeft, int frontRight,
                                 int rearLeft, int rearRight) {
  return api.Vehicle.SetParamsCarWindows(vehicle, frontLeft, frontRight,
                                         rearLeft, rearRight);
}

bool Vehicle_GetParamsCarWindows(void *vehicle, int *frontLeft, int *frontRight,
                                 int *rearLeft, int *rearRight) {
  return api.Vehicle.GetParamsCarWindows(vehicle, frontLeft, frontRight,
                                         rearLeft, rearRight);
}

bool Vehicle_SetToRespawn(void *vehicle) {
  return api.Vehicle.SetToRespawn(vehicle);
}

bool Vehicle_LinkToInterior(void *vehicle, int interiorid) {
  return api.Vehicle.LinkToInterior(vehicle, interiorid);
}

bool Vehicle_AddComponent(void *vehicle, int componentid) {
  return api.Vehicle.AddComponent(vehicle, componentid);
}

bool Vehicle_RemoveComponent(void *vehicle, int componentid) {
  return api.Vehicle.RemoveComponent(vehicle, componentid);
}

bool Vehicle_ChangeColor(void *vehicle, int color1, int color2) {
  return api.Vehicle.ChangeColor(vehicle, color1, color2);
}

bool Vehicle_ChangePaintjob(void *vehicle, int paintjobid) {
  return api.Vehicle.ChangePaintjob(vehicle, paintjobid);
}

bool Vehicle_SetHealth(void *vehicle, float health) {
  return api.Vehicle.SetHealth(vehicle, health);
}

float Vehicle_GetHealth(void *vehicle) {
  return api.Vehicle.GetHealth(vehicle);
}

bool Vehicle_AttachTrailer(void *trailer, void *vehicle) {
  return api.Vehicle.AttachTrailer(trailer, vehicle);
}

bool Vehicle_DetachTrailer(void *vehicle) {
  return api.Vehicle.DetachTrailer(vehicle);
}

bool Vehicle_IsTrailerAttached(void *vehicle) {
  return api.Vehicle.IsTrailerAttached(vehicle);
}

void *Vehicle_GetTrailer(void *vehicle) {
  return api.Vehicle.GetTrailer(vehicle);
}

bool Vehicle_SetNumberPlate(void *vehicle, const char *numberPlate) {
  return api.Vehicle.SetNumberPlate(vehicle, numberPlate);
}

int Vehicle_GetModel(void *vehicle) { return api.Vehicle.GetModel(vehicle); }

int Vehicle_GetComponentInSlot(void *vehicle, int slot) {
  return api.Vehicle.GetComponentInSlot(vehicle, slot);
}

int Vehicle_GetComponentType(int componentid) {
  return api.Vehicle.GetComponentType(componentid);
}

bool Vehicle_CanHaveComponent(int modelid, int componentid) {
  return api.Vehicle.CanHaveComponent(modelid, componentid);
}

bool Vehicle_GetRandomColorPair(int modelid, int *color1, int *color2,
                                int *color3, int *color4) {
  return api.Vehicle.GetRandomColorPair(modelid, color1, color2, color3,
                                        color4);
}

int Vehicle_ColorIndexToColor(int colorIndex, int alpha) {
  return api.Vehicle.ColorIndexToColor(colorIndex, alpha);
}

bool Vehicle_Repair(void *vehicle) { return api.Vehicle.Repair(vehicle); }

bool Vehicle_GetVelocity(void *vehicle, float *x, float *y, float *z) {
  return api.Vehicle.GetVelocity(vehicle, x, y, z);
}

bool Vehicle_SetVelocity(void *vehicle, float x, float y, float z) {
  return api.Vehicle.SetVelocity(vehicle, x, y, z);
}

bool Vehicle_SetAngularVelocity(void *vehicle, float x, float y, float z) {
  return api.Vehicle.SetAngularVelocity(vehicle, x, y, z);
}

bool Vehicle_GetDamageStatus(void *vehicle, int *panels, int *doors,
                             int *lights, int *tires) {
  return api.Vehicle.GetDamageStatus(vehicle, panels, doors, lights, tires);
}

bool Vehicle_UpdateDamageStatus(void *vehicle, int panels, int doors,
                                int lights, int tires) {
  return api.Vehicle.UpdateDamageStatus(vehicle, panels, doors, lights, tires);
}

bool Vehicle_GetModelInfo(int vehiclemodel, int infotype, float *x, float *y,
                          float *z) {
  return api.Vehicle.GetModelInfo(vehiclemodel, infotype, x, y, z);
}

bool Vehicle_SetVirtualWorld(void *vehicle, int virtualWorld) {
  return api.Vehicle.SetVirtualWorld(vehicle, virtualWorld);
}

int Vehicle_GetVirtualWorld(void *vehicle) {
  return api.Vehicle.GetVirtualWorld(vehicle);
}

int Vehicle_GetLandingGearState(void *vehicle) {
  return api.Vehicle.GetLandingGearState(vehicle);
}

bool Vehicle_IsValid(void *vehicle) { return api.Vehicle.IsValid(vehicle); }

void *Vehicle_AddStatic(int modelid, float x, float y, float z, float angle,
                        int color1, int color2, int *id) {
  return api.Vehicle.AddStatic(modelid, x, y, z, angle, color1, color2, id);
}

void *Vehicle_AddStaticEx(int modelid, float x, float y, float z, float angle,
                          int color1, int color2, int respawnDelay,
                          bool addSiren, int *id) {
  return api.Vehicle.AddStaticEx(modelid, x, y, z, angle, color1, color2,
                                 respawnDelay, addSiren, id);
}

bool Vehicle_EnableFriendlyFire() { return api.Vehicle.EnableFriendlyFire(); }

bool Vehicle_GetSpawnInfo(void *vehicle, float *x, float *y, float *z,
                          float *rotation, int *color1, int *color2) {
  return api.Vehicle.GetSpawnInfo(vehicle, x, y, z, rotation, color1, color2);
}

bool Vehicle_SetSpawnInfo(void *vehicle, int modelid, float x, float y, float z,
                          float rotation, int color1, int color2,
                          int respawn_time, int interior) {
  return api.Vehicle.SetSpawnInfo(vehicle, modelid, x, y, z, rotation, color1,
                                  color2, respawn_time, interior);
}

int Vehicle_GetModelCount(int modelid) {
  return api.Vehicle.GetModelCount(modelid);
}

int Vehicle_GetModelsUsed() { return api.Vehicle.GetModelsUsed(); }

int Vehicle_GetPaintjob(void *vehicle) {
  return api.Vehicle.GetPaintjob(vehicle);
}

bool Vehicle_GetColor(void *vehicle, int *color1, int *color2) {
  return api.Vehicle.GetColor(vehicle, color1, color2);
}

int Vehicle_GetInterior(void *vehicle) {
  return api.Vehicle.GetInterior(vehicle);
}

bool Vehicle_GetNumberPlate(void *vehicle, struct CAPIStringView *numberPlate) {
  return api.Vehicle.GetNumberPlate(vehicle, numberPlate);
}

bool Vehicle_SetRespawnDelay(void *vehicle, int respawn_delay) {
  return api.Vehicle.SetRespawnDelay(vehicle, respawn_delay);
}

int Vehicle_GetRespawnDelay(void *vehicle) {
  return api.Vehicle.GetRespawnDelay(vehicle);
}

void *Vehicle_GetCab(void *vehicle) { return api.Vehicle.GetCab(vehicle); }

void *Vehicle_GetTower(void *vehicle) { return api.Vehicle.GetTower(vehicle); }

int Vehicle_GetOccupiedTick(void *vehicle) {
  return api.Vehicle.GetOccupiedTick(vehicle);
}

int Vehicle_GetRespawnTick(void *vehicle) {
  return api.Vehicle.GetRespawnTick(vehicle);
}

bool Vehicle_HasBeenOccupied(void *vehicle) {
  return api.Vehicle.HasBeenOccupied(vehicle);
}

bool Vehicle_IsOccupied(void *vehicle) {
  return api.Vehicle.IsOccupied(vehicle);
}

bool Vehicle_IsDead(void *vehicle) { return api.Vehicle.IsDead(vehicle); }

bool Vehicle_SetParamsSirenState(void *vehicle, bool siren_state) {
  return api.Vehicle.SetParamsSirenState(vehicle, siren_state);
}

bool Vehicle_ToggleSirenEnabled(void *vehicle, bool status) {
  return api.Vehicle.ToggleSirenEnabled(vehicle, status);
}

bool Vehicle_IsSirenEnabled(void *vehicle) {
  return api.Vehicle.IsSirenEnabled(vehicle);
}

void *Vehicle_GetLastDriver(void *vehicle) {
  return api.Vehicle.GetLastDriver(vehicle);
}

void *Vehicle_GetDriver(void *vehicle) {
  return api.Vehicle.GetDriver(vehicle);
}

int Vehicle_GetSirenState(void *vehicle) {
  return api.Vehicle.GetSirenState(vehicle);
}

uint32_t Vehicle_GetHydraReactorAngle(void *vehicle) {
  return api.Vehicle.GetHydraReactorAngle(vehicle);
}

float Vehicle_GetTrainSpeed(void *vehicle) {
  return api.Vehicle.GetTrainSpeed(vehicle);
}

bool Vehicle_GetMatrix(void *vehicle, float *rightX, float *rightY,
                       float *rightZ, float *upX, float *upY, float *upZ,
                       float *atX, float *atY, float *atZ) {
  return api.Vehicle.GetMatrix(vehicle, rightX, rightY, rightZ, upX, upY, upZ,
                               atX, atY, atZ);
}

void *Vehicle_GetOccupant(void *vehicle, int seat) {
  return api.Vehicle.GetOccupant(vehicle, seat);
}

int Vehicle_CountOccupants(void *vehicle) {
  return api.Vehicle.CountOccupants(vehicle);
}
