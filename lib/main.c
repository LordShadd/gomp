#include "main.h"
#include "ompcapi.h"

struct OMPAPI_t api;

struct ComponentVersion componentVersion;
char *componentName = NULL;

void setComponentVersion(uint8_t major, uint8_t minor, uint8_t patch,
                         uint16_t prerel) {
  componentVersion = (struct ComponentVersion){major, minor, patch, prerel};
}

void setComponentName(char *name) {
  if (componentName != NULL) {
    free(componentName);
  }

  componentName = strdup(name);
}

void _onReady() { printf("onReady\n"); }
void _onReset() { printf("onReset\n"); }
void _onFree() {
  printf("onFree\n");

  if (componentName != NULL) {
    free(componentName);
  }
}

OMP_API_EXPORT void *ComponentEntryPoint() {
  if (!omp_initialize_capi(&api)) {
    printf("Failed to initialize open.mp C API\n");
    return NULL;
  }

  entryPoint();

  if (componentName == NULL) {
    componentName = strdup("GOMP GAMEMODE");
  }

  void *comp =
      api.Component.Create(0x913B89092F8F6A68, componentName, componentVersion,
                           &_onReady, &_onReset, &_onFree);

  return comp;
}
