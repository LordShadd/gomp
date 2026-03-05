#include "playertextdraw.h"
#include "ompcapi.h"

void *PlayerTextDraw_Create(void *player, float x, float y, const char *text,
                             int *id) {
  return api.PlayerTextDraw.Create(player, x, y, text, id);
}

bool PlayerTextDraw_Destroy(void *player, void *textdraw) {
  return api.PlayerTextDraw.Destroy(player, textdraw);
}

void *PlayerTextDraw_FromID(void *player, int textdrawid) {
  return api.PlayerTextDraw.FromID(player, textdrawid);
}

int PlayerTextDraw_GetID(void *player, void *textdraw) {
  return api.PlayerTextDraw.GetID(player, textdraw);
}

bool PlayerTextDraw_IsValid(void *player, void *textdraw) {
  return api.PlayerTextDraw.IsValid(player, textdraw);
}

bool PlayerTextDraw_IsVisible(void *player, void *textdraw) {
  return api.PlayerTextDraw.IsVisible(player, textdraw);
}

bool PlayerTextDraw_SetLetterSize(void *player, void *textdraw, float x,
                                  float y) {
  return api.PlayerTextDraw.SetLetterSize(player, textdraw, x, y);
}

bool PlayerTextDraw_SetTextSize(void *player, void *textdraw, float x,
                                float y) {
  return api.PlayerTextDraw.SetTextSize(player, textdraw, x, y);
}

bool PlayerTextDraw_SetAlignment(void *player, void *textdraw, int alignment) {
  return api.PlayerTextDraw.SetAlignment(player, textdraw, alignment);
}

bool PlayerTextDraw_SetColor(void *player, void *textdraw, uint32_t color) {
  return api.PlayerTextDraw.SetColor(player, textdraw, color);
}

bool PlayerTextDraw_UseBox(void *player, void *textdraw, bool use) {
  return api.PlayerTextDraw.UseBox(player, textdraw, use);
}

bool PlayerTextDraw_SetBoxColor(void *player, void *textdraw, uint32_t color) {
  return api.PlayerTextDraw.SetBoxColor(player, textdraw, color);
}

bool PlayerTextDraw_SetShadow(void *player, void *textdraw, int size) {
  return api.PlayerTextDraw.SetShadow(player, textdraw, size);
}

bool PlayerTextDraw_SetOutline(void *player, void *textdraw, int size) {
  return api.PlayerTextDraw.SetOutline(player, textdraw, size);
}

bool PlayerTextDraw_SetBackgroundColor(void *player, void *textdraw,
                                       uint32_t color) {
  return api.PlayerTextDraw.SetBackgroundColor(player, textdraw, color);
}

bool PlayerTextDraw_SetFont(void *player, void *textdraw, int font) {
  return api.PlayerTextDraw.SetFont(player, textdraw, font);
}

bool PlayerTextDraw_SetProportional(void *player, void *textdraw, bool set) {
  return api.PlayerTextDraw.SetProportional(player, textdraw, set);
}

bool PlayerTextDraw_SetSelectable(void *player, void *textdraw, bool set) {
  return api.PlayerTextDraw.SetSelectable(player, textdraw, set);
}

bool PlayerTextDraw_Show(void *player, void *textdraw) {
  return api.PlayerTextDraw.Show(player, textdraw);
}

bool PlayerTextDraw_Hide(void *player, void *textdraw) {
  return api.PlayerTextDraw.Hide(player, textdraw);
}

bool PlayerTextDraw_SetString(void *player, void *textdraw, const char *text) {
  return api.PlayerTextDraw.SetString(player, textdraw, text);
}

bool PlayerTextDraw_SetPreviewModel(void *player, void *textdraw, int model) {
  return api.PlayerTextDraw.SetPreviewModel(player, textdraw, model);
}

bool PlayerTextDraw_SetPreviewRot(void *player, void *textdraw, float rx,
                                  float ry, float rz, float zoom) {
  return api.PlayerTextDraw.SetPreviewRot(player, textdraw, rx, ry, rz, zoom);
}

bool PlayerTextDraw_SetPreviewVehCol(void *player, void *textdraw, int color1,
                                     int color2) {
  return api.PlayerTextDraw.SetPreviewVehCol(player, textdraw, color1, color2);
}

bool PlayerTextDraw_SetPos(void *player, void *textdraw, float x, float y) {
  return api.PlayerTextDraw.SetPos(player, textdraw, x, y);
}

bool PlayerTextDraw_GetString(void *player, void *textdraw,
                              struct CAPIStringView *text) {
  return api.PlayerTextDraw.GetString(player, textdraw, text);
}

bool PlayerTextDraw_GetLetterSize(void *player, void *textdraw, float *x,
                                  float *y) {
  return api.PlayerTextDraw.GetLetterSize(player, textdraw, x, y);
}

bool PlayerTextDraw_GetTextSize(void *player, void *textdraw, float *x,
                                float *y) {
  return api.PlayerTextDraw.GetTextSize(player, textdraw, x, y);
}

bool PlayerTextDraw_GetPos(void *player, void *textdraw, float *x, float *y) {
  return api.PlayerTextDraw.GetPos(player, textdraw, x, y);
}

int PlayerTextDraw_GetColor(void *player, void *textdraw) {
  return api.PlayerTextDraw.GetColor(player, textdraw);
}

int PlayerTextDraw_GetBoxColor(void *player, void *textdraw) {
  return api.PlayerTextDraw.GetBoxColor(player, textdraw);
}

int PlayerTextDraw_GetBackgroundColor(void *player, void *textdraw) {
  return api.PlayerTextDraw.GetBackgroundColor(player, textdraw);
}

int PlayerTextDraw_GetShadow(void *player, void *textdraw) {
  return api.PlayerTextDraw.GetShadow(player, textdraw);
}

int PlayerTextDraw_GetOutline(void *player, void *textdraw) {
  return api.PlayerTextDraw.GetOutline(player, textdraw);
}

int PlayerTextDraw_GetFont(void *player, void *textdraw) {
  return api.PlayerTextDraw.GetFont(player, textdraw);
}

bool PlayerTextDraw_IsBox(void *player, void *textdraw) {
  return api.PlayerTextDraw.IsBox(player, textdraw);
}

bool PlayerTextDraw_IsProportional(void *player, void *textdraw) {
  return api.PlayerTextDraw.IsProportional(player, textdraw);
}

bool PlayerTextDraw_IsSelectable(void *player, void *textdraw) {
  return api.PlayerTextDraw.IsSelectable(player, textdraw);
}

int PlayerTextDraw_GetAlignment(void *player, void *textdraw) {
  return api.PlayerTextDraw.GetAlignment(player, textdraw);
}

int PlayerTextDraw_GetPreviewModel(void *player, void *textdraw) {
  return api.PlayerTextDraw.GetPreviewModel(player, textdraw);
}

bool PlayerTextDraw_GetPreviewRot(void *player, void *textdraw, float *rx,
                                  float *ry, float *rz, float *zoom) {
  return api.PlayerTextDraw.GetPreviewRot(player, textdraw, rx, ry, rz, zoom);
}

bool PlayerTextDraw_GetPreviewVehColor(void *player, void *textdraw,
                                       int *color1, int *color2) {
  return api.PlayerTextDraw.GetPreviewVehColor(player, textdraw, color1,
                                               color2);
}
