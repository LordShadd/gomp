#ifndef VEHICLE_H
#define VEHICLE_H

#include <stdbool.h>
#include <stdint.h>

void *Vehicle_Create(int modelid, float x, float y, float z, float rotation,
                     int color1, int color2, int respawnDelay, bool addSiren,
                     int *id);
bool Vehicle_Destroy(void *vehicle);
void *Vehicle_FromID(int vehicleid);
int Vehicle_GetID(void *vehicle);
int Vehicle_GetMaxPassengerSeats(int modelid);
bool Vehicle_IsStreamedIn(void *vehicle, void *player);
bool Vehicle_GetPos(void *vehicle, float *x, float *y, float *z);
bool Vehicle_SetPos(void *vehicle, float x, float y, float z);
float Vehicle_GetZAngle(void *vehicle);
bool Vehicle_GetRotationQuat(void *vehicle, float *w, float *x, float *y,
                             float *z);
float Vehicle_GetDistanceFromPoint(void *vehicle, float x, float y, float z);
bool Vehicle_SetZAngle(void *vehicle, float angle);
bool Vehicle_SetParamsForPlayer(void *vehicle, void *player, int objective,
                                int doors);
bool Vehicle_UseManualEngineAndLights();
bool Vehicle_SetParamsEx(void *vehicle, int engine, int lights, int alarm,
                         int doors, int bonnet, int boot, int objective);
bool Vehicle_GetParamsEx(void *vehicle, int *engine, int *lights, int *alarm,
                         int *doors, int *bonnet, int *boot, int *objective);
int Vehicle_GetParamsSirenState(void *vehicle);
bool Vehicle_SetParamsCarDoors(void *vehicle, int frontLeft, int frontRight,
                               int rearLeft, int rearRight);
bool Vehicle_GetParamsCarDoors(void *vehicle, int *frontLeft, int *frontRight,
                               int *rearLeft, int *rearRight);
bool Vehicle_SetParamsCarWindows(void *vehicle, int frontLeft, int frontRight,
                                 int rearLeft, int rearRight);
bool Vehicle_GetParamsCarWindows(void *vehicle, int *frontLeft, int *frontRight,
                                 int *rearLeft, int *rearRight);
bool Vehicle_SetToRespawn(void *vehicle);
bool Vehicle_LinkToInterior(void *vehicle, int interiorid);
bool Vehicle_AddComponent(void *vehicle, int componentid);
bool Vehicle_RemoveComponent(void *vehicle, int componentid);
bool Vehicle_ChangeColor(void *vehicle, int color1, int color2);
bool Vehicle_ChangePaintjob(void *vehicle, int paintjobid);
bool Vehicle_SetHealth(void *vehicle, float health);
float Vehicle_GetHealth(void *vehicle);
bool Vehicle_AttachTrailer(void *trailer, void *vehicle);
bool Vehicle_DetachTrailer(void *vehicle);
bool Vehicle_IsTrailerAttached(void *vehicle);
void *Vehicle_GetTrailer(void *vehicle);
bool Vehicle_SetNumberPlate(void *vehicle, const char *numberPlate);
int Vehicle_GetModel(void *vehicle);
int Vehicle_GetComponentInSlot(void *vehicle, int slot);
int Vehicle_GetComponentType(int componentid);
bool Vehicle_CanHaveComponent(int modelid, int componentid);
bool Vehicle_GetRandomColorPair(int modelid, int *color1, int *color2,
                                int *color3, int *color4);
int Vehicle_ColorIndexToColor(int colorIndex, int alpha);
bool Vehicle_Repair(void *vehicle);
bool Vehicle_GetVelocity(void *vehicle, float *x, float *y, float *z);
bool Vehicle_SetVelocity(void *vehicle, float x, float y, float z);
bool Vehicle_SetAngularVelocity(void *vehicle, float x, float y, float z);
bool Vehicle_GetDamageStatus(void *vehicle, int *panels, int *doors,
                             int *lights, int *tires);
bool Vehicle_UpdateDamageStatus(void *vehicle, int panels, int doors,
                                int lights, int tires);
bool Vehicle_GetModelInfo(int vehiclemodel, int infotype, float *x, float *y,
                          float *z);
bool Vehicle_SetVirtualWorld(void *vehicle, int virtualWorld);
int Vehicle_GetVirtualWorld(void *vehicle);
int Vehicle_GetLandingGearState(void *vehicle);
bool Vehicle_IsValid(void *vehicle);
void *Vehicle_AddStatic(int modelid, float x, float y, float z, float angle,
                        int color1, int color2, int *id);
void *Vehicle_AddStaticEx(int modelid, float x, float y, float z, float angle,
                          int color1, int color2, int respawnDelay,
                          bool addSiren, int *id);
bool Vehicle_EnableFriendlyFire();
bool Vehicle_GetSpawnInfo(void *vehicle, float *x, float *y, float *z,
                          float *rotation, int *color1, int *color2);
bool Vehicle_SetSpawnInfo(void *vehicle, int modelid, float x, float y, float z,
                          float rotation, int color1, int color2,
                          int respawn_time, int interior);
int Vehicle_GetModelCount(int modelid);
int Vehicle_GetModelsUsed();
int Vehicle_GetPaintjob(void *vehicle);
bool Vehicle_GetColor(void *vehicle, int *color1, int *color2);
int Vehicle_GetInterior(void *vehicle);
bool Vehicle_GetNumberPlate(void *vehicle, struct CAPIStringView *numberPlate);
bool Vehicle_SetRespawnDelay(void *vehicle, int respawn_delay);
int Vehicle_GetRespawnDelay(void *vehicle);
void *Vehicle_GetCab(void *vehicle);
void *Vehicle_GetTower(void *vehicle);
int Vehicle_GetOccupiedTick(void *vehicle);
int Vehicle_GetRespawnTick(void *vehicle);
bool Vehicle_HasBeenOccupied(void *vehicle);
bool Vehicle_IsOccupied(void *vehicle);
bool Vehicle_IsDead(void *vehicle);
bool Vehicle_SetParamsSirenState(void *vehicle, bool siren_state);
bool Vehicle_ToggleSirenEnabled(void *vehicle, bool status);
bool Vehicle_IsSirenEnabled(void *vehicle);
void *Vehicle_GetLastDriver(void *vehicle);
void *Vehicle_GetDriver(void *vehicle);
int Vehicle_GetSirenState(void *vehicle);
uint32_t Vehicle_GetHydraReactorAngle(void *vehicle);
float Vehicle_GetTrainSpeed(void *vehicle);
bool Vehicle_GetMatrix(void *vehicle, float *rightX, float *rightY,
                       float *rightZ, float *upX, float *upY, float *upZ,
                       float *atX, float *atY, float *atZ);
void *Vehicle_GetOccupant(void *vehicle, int seat);
int Vehicle_CountOccupants(void *vehicle);

#endif
