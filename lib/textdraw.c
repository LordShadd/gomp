#include "textdraw.h"
#include "ompcapi.h"

void *TextDraw_Create(float x, float y, const char *text, int *id) {
  return api.TextDraw.Create(x, y, text, id);
}

bool TextDraw_Destroy(void *textdraw) {
  return api.TextDraw.Destroy(textdraw);
}

void *TextDraw_FromID(int textdrawid) {
  return api.TextDraw.FromID(textdrawid);
}

int TextDraw_GetID(void *textdraw) { return api.TextDraw.GetID(textdraw); }

bool TextDraw_IsValid(void *textdraw) {
  return api.TextDraw.IsValid(textdraw);
}

bool TextDraw_IsVisibleForPlayer(void *player, void *textdraw) {
  return api.TextDraw.IsVisibleForPlayer(player, textdraw);
}

bool TextDraw_SetLetterSize(void *textdraw, float sizeX, float sizeY) {
  return api.TextDraw.SetLetterSize(textdraw, sizeX, sizeY);
}

bool TextDraw_SetTextSize(void *textdraw, float sizeX, float sizeY) {
  return api.TextDraw.SetTextSize(textdraw, sizeX, sizeY);
}

bool TextDraw_SetAlignment(void *textdraw, int alignment) {
  return api.TextDraw.SetAlignment(textdraw, alignment);
}

bool TextDraw_SetColor(void *textdraw, uint32_t color) {
  return api.TextDraw.SetColor(textdraw, color);
}

bool TextDraw_SetUseBox(void *textdraw, bool use) {
  return api.TextDraw.SetUseBox(textdraw, use);
}

bool TextDraw_SetBoxColor(void *textdraw, uint32_t color) {
  return api.TextDraw.SetBoxColor(textdraw, color);
}

bool TextDraw_SetShadow(void *textdraw, int size) {
  return api.TextDraw.SetShadow(textdraw, size);
}

bool TextDraw_SetOutline(void *textdraw, int size) {
  return api.TextDraw.SetOutline(textdraw, size);
}

bool TextDraw_SetBackgroundColor(void *textdraw, uint32_t color) {
  return api.TextDraw.SetBackgroundColor(textdraw, color);
}

bool TextDraw_SetFont(void *textdraw, int font) {
  return api.TextDraw.SetFont(textdraw, font);
}

bool TextDraw_SetProportional(void *textdraw, bool set) {
  return api.TextDraw.SetProportional(textdraw, set);
}

bool TextDraw_SetSelectable(void *textdraw, bool set) {
  return api.TextDraw.SetSelectable(textdraw, set);
}

bool TextDraw_ShowForPlayer(void *player, void *textdraw) {
  return api.TextDraw.ShowForPlayer(player, textdraw);
}

bool TextDraw_HideForPlayer(void *player, void *textdraw) {
  return api.TextDraw.HideForPlayer(player, textdraw);
}

bool TextDraw_ShowForAll(void *textdraw) {
  return api.TextDraw.ShowForAll(textdraw);
}

bool TextDraw_HideForAll(void *textdraw) {
  return api.TextDraw.HideForAll(textdraw);
}

bool TextDraw_SetString(void *textdraw, const char *text) {
  return api.TextDraw.SetString(textdraw, text);
}

bool TextDraw_SetPreviewModel(void *textdraw, int model) {
  return api.TextDraw.SetPreviewModel(textdraw, model);
}

bool TextDraw_SetPreviewRot(void *textdraw, float rotationX, float rotationY,
                            float rotationZ, float zoom) {
  return api.TextDraw.SetPreviewRot(textdraw, rotationX, rotationY, rotationZ,
                                    zoom);
}

bool TextDraw_SetPreviewVehCol(void *textdraw, int color1, int color2) {
  return api.TextDraw.SetPreviewVehCol(textdraw, color1, color2);
}

bool TextDraw_SetPos(void *textdraw, float x, float y) {
  return api.TextDraw.SetPos(textdraw, x, y);
}

bool TextDraw_GetString(void *textdraw, struct CAPIStringView *text) {
  return api.TextDraw.GetString(textdraw, text);
}

bool TextDraw_GetLetterSize(void *textdraw, float *sizeX, float *sizeY) {
  return api.TextDraw.GetLetterSize(textdraw, sizeX, sizeY);
}

bool TextDraw_GetTextSize(void *textdraw, float *sizeX, float *sizeY) {
  return api.TextDraw.GetTextSize(textdraw, sizeX, sizeY);
}

bool TextDraw_GetPos(void *textdraw, float *x, float *y) {
  return api.TextDraw.GetPos(textdraw, x, y);
}

int TextDraw_GetColor(void *textdraw) {
  return api.TextDraw.GetColor(textdraw);
}

int TextDraw_GetBoxColor(void *textdraw) {
  return api.TextDraw.GetBoxColor(textdraw);
}

int TextDraw_GetBackgroundColor(void *textdraw) {
  return api.TextDraw.GetBackgroundColor(textdraw);
}

int TextDraw_GetShadow(void *textdraw) {
  return api.TextDraw.GetShadow(textdraw);
}

int TextDraw_GetOutline(void *textdraw) {
  return api.TextDraw.GetOutline(textdraw);
}

int TextDraw_GetFont(void *textdraw) { return api.TextDraw.GetFont(textdraw); }

bool TextDraw_IsBox(void *textdraw) { return api.TextDraw.IsBox(textdraw); }

bool TextDraw_IsProportional(void *textdraw) {
  return api.TextDraw.IsProportional(textdraw);
}

bool TextDraw_IsSelectable(void *textdraw) {
  return api.TextDraw.IsSelectable(textdraw);
}

int TextDraw_GetAlignment(void *textdraw) {
  return api.TextDraw.GetAlignment(textdraw);
}

int TextDraw_GetPreviewModel(void *textdraw) {
  return api.TextDraw.GetPreviewModel(textdraw);
}

bool TextDraw_GetPreviewRot(void *textdraw, float *x, float *y, float *z,
                            float *zoom) {
  return api.TextDraw.GetPreviewRot(textdraw, x, y, z, zoom);
}

bool TextDraw_GetPreviewVehColor(void *textdraw, int *color1, int *color2) {
  return api.TextDraw.GetPreviewVehColor(textdraw, color1, color2);
}

bool TextDraw_SetStringForPlayer(void *textdraw, void *player,
                                 const char *text) {
  return api.TextDraw.SetStringForPlayer(textdraw, player, text);
}
