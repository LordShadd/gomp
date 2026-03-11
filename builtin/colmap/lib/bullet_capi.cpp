#include "bullet_capi.h"
#include <bullet/btBulletDynamicsCommon.h>
#include <cmath>
#include <unordered_map>

struct CMWorldImpl {
    btBroadphaseInterface*               broadphase;
    btDefaultCollisionConfiguration*     collisionConfig;
    btCollisionDispatcher*               dispatcher;
    btSequentialImpulseConstraintSolver* solver;
    btDiscreteDynamicsWorld*             dynamicsWorld;

    std::unordered_map<int, btRigidBody*> bodies;
    int nextHandle = 0;

    CMWorldImpl() {
        broadphase      = new btDbvtBroadphase();
        collisionConfig = new btDefaultCollisionConfiguration();
        dispatcher      = new btCollisionDispatcher(collisionConfig);
        solver          = new btSequentialImpulseConstraintSolver();
        dynamicsWorld   = new btDiscreteDynamicsWorld(dispatcher, broadphase, solver, collisionConfig);
    }

    ~CMWorldImpl() {
        for (auto &kv : bodies) {
            dynamicsWorld->removeRigidBody(kv.second);
            delete kv.second->getMotionState();
            delete kv.second;
        }
        delete dynamicsWorld;
        delete solver;
        delete dispatcher;
        delete collisionConfig;
        delete broadphase;
    }
};

struct CMShapeImpl {
    btCompoundShape*              compound;
    btTriangleMesh*               trimesh;
    btBvhTriangleMeshShape*       meshShape;
    btConvexTriangleMeshShape*    convexShape;

    CMShapeImpl() : compound(nullptr), trimesh(nullptr), meshShape(nullptr), convexShape(nullptr) {
        compound = new btCompoundShape();
    }

    ~CMShapeImpl() {
        delete compound;
        delete meshShape;
        delete convexShape;
        delete trimesh;
    }
};

struct ContactSensor : public btCollisionWorld::ContactResultCallback {
    int collided = 0;
    virtual btScalar addSingleResult(btManifoldPoint&,
                                      const btCollisionObjectWrapper*, int, int,
                                      const btCollisionObjectWrapper*, int, int) override {
        collided = 1;
        return 0;
    }
};

extern "C" {

CMWorld cm_world_create() {
    return new CMWorldImpl();
}

void cm_world_destroy(CMWorld w) {
    if (w) delete static_cast<CMWorldImpl*>(w);
}

CMShape cm_shape_create() {
    return new CMShapeImpl();
}

void cm_shape_destroy(CMShape s) {
    if (s) delete static_cast<CMShapeImpl*>(s);
}

void cm_shape_add_sphere(CMShape s, float offX, float offY, float offZ, float radius) {
    if (!s) return;
    auto *impl = static_cast<CMShapeImpl*>(s);
    btSphereShape *sphere = new btSphereShape(btScalar(radius));
    impl->compound->addChildShape(
        btTransform(btQuaternion(0, 0, 0, 1), btVector3(offX, offY, offZ)),
        sphere);
}

void cm_shape_add_box(CMShape s, float cx, float cy, float cz, float sx, float sy, float sz) {
    if (!s) return;
    auto *impl = static_cast<CMShapeImpl*>(s);
    btBoxShape *box = new btBoxShape(btVector3(sx, sy, sz));
    impl->compound->addChildShape(
        btTransform(btQuaternion(0, 0, 0, 1), btVector3(cx, cy, cz)),
        box);
}

void cm_shape_add_trimesh(CMShape s, float *verts, int numVerts, int *indices, int numTris, int convex) {
    if (!s || !verts || !indices) return;
    auto *impl = static_cast<CMShapeImpl*>(s);

    if (!impl->trimesh) impl->trimesh = new btTriangleMesh();
    for (int i = 0; i < numTris; i++) {
        int ia = indices[i * 3];
        int ib = indices[i * 3 + 1];
        int ic = indices[i * 3 + 2];
        impl->trimesh->addTriangle(
            btVector3(verts[ia*3], verts[ia*3+1], verts[ia*3+2]),
            btVector3(verts[ib*3], verts[ib*3+1], verts[ib*3+2]),
            btVector3(verts[ic*3], verts[ic*3+1], verts[ic*3+2]));
    }

    if (convex) {
        impl->convexShape = new btConvexTriangleMeshShape(impl->trimesh);
        impl->compound->addChildShape(
            btTransform(btQuaternion(0, 0, 0, 1), btVector3(0, 0, 0)),
            impl->convexShape);
    } else {
        impl->meshShape = new btBvhTriangleMeshShape(impl->trimesh, true);
        impl->compound->addChildShape(
            btTransform(btQuaternion(0, 0, 0, 1), btVector3(0, 0, 0)),
            impl->meshShape);
    }
}

void cm_shape_bounding_sphere(CMShape s, float *cx, float *cy, float *cz, float *r) {
    if (!s) return;
    auto *impl = static_cast<CMShapeImpl*>(s);
    btVector3 center;
    btScalar radius;
    impl->compound->getBoundingSphere(center, radius);
    if (cx) *cx = center.getX();
    if (cy) *cy = center.getY();
    if (cz) *cz = center.getZ();
    if (r)  *r  = radius;
}

void cm_shape_bounding_box(CMShape s, float *minx, float *miny, float *minz,
                            float *maxx, float *maxy, float *maxz) {
    if (!s) return;
    auto *impl = static_cast<CMShapeImpl*>(s);
    btTransform t;
    t.setIdentity();
    btVector3 bmin, bmax;
    impl->compound->getAabb(t, bmin, bmax);
    if (minx) *minx = bmin.getX();
    if (miny) *miny = bmin.getY();
    if (minz) *minz = bmin.getZ();
    if (maxx) *maxx = bmax.getX();
    if (maxy) *maxy = bmax.getY();
    if (maxz) *maxz = bmax.getZ();
}

int cm_world_add_body(CMWorld w, CMShape s,
                       float x, float y, float z,
                       float qx, float qy, float qz, float qw,
                       int modelId, int bodyTag) {
    if (!w || !s) return -1;
    auto *world  = static_cast<CMWorldImpl*>(w);
    auto *shape  = static_cast<CMShapeImpl*>(s);
    btDefaultMotionState *ms = new btDefaultMotionState(
        btTransform(btQuaternion(qx, qy, qz, qw), btVector3(x, y, z)));
    btRigidBody::btRigidBodyConstructionInfo ci(0, ms, shape->compound, btVector3(0, 0, 0));
    btRigidBody *body = new btRigidBody(ci);
    body->setUserIndex(modelId);
    body->setUserIndex2(bodyTag);
    world->dynamicsWorld->addRigidBody(body);
    int handle = world->nextHandle++;
    world->bodies[handle] = body;
    return handle;
}

void cm_world_remove_body(CMWorld w, int handle) {
    if (!w) return;
    auto *world = static_cast<CMWorldImpl*>(w);
    auto it = world->bodies.find(handle);
    if (it == world->bodies.end()) return;
    world->dynamicsWorld->removeRigidBody(it->second);
    delete it->second->getMotionState();
    delete it->second;
    world->bodies.erase(it);
}

int cm_world_set_body_pos(CMWorld w, int handle, float x, float y, float z) {
    if (!w) return 0;
    auto *world = static_cast<CMWorldImpl*>(w);
    auto it = world->bodies.find(handle);
    if (it == world->bodies.end()) return 0;
    btRigidBody *body = it->second;
    btTransform t = body->getWorldTransform();
    t.setOrigin(btVector3(x, y, z));
    body->setWorldTransform(t);
    world->dynamicsWorld->removeRigidBody(body);
    world->dynamicsWorld->addRigidBody(body);
    return 1;
}

int cm_world_set_body_rot(CMWorld w, int handle, float qx, float qy, float qz, float qw) {
    if (!w) return 0;
    auto *world = static_cast<CMWorldImpl*>(w);
    auto it = world->bodies.find(handle);
    if (it == world->bodies.end()) return 0;
    btRigidBody *body = it->second;
    btTransform t = body->getWorldTransform();
    t.setRotation(btQuaternion(qx, qy, qz, qw));
    body->setWorldTransform(t);
    world->dynamicsWorld->removeRigidBody(body);
    world->dynamicsWorld->addRigidBody(body);
    return 1;
}

int cm_raytest(CMWorld w,
               float sx, float sy, float sz,
               float ex, float ey, float ez,
               float *rx, float *ry, float *rz,
               float *nx, float *ny, float *nz,
               int *modelId, int *bodyTag) {
    if (!w) return 0;
    auto *world = static_cast<CMWorldImpl*>(w);
    btVector3 Start(sx + 0.00001f, sy + 0.00001f, sz + 0.00001f);
    btVector3 End(ex, ey, ez);
    btCollisionWorld::ClosestRayResultCallback cb(Start, End);
    world->dynamicsWorld->rayTest(Start, End, cb);
    if (!cb.hasHit()) return 0;
    if (rx) *rx = cb.m_hitPointWorld.getX();
    if (ry) *ry = cb.m_hitPointWorld.getY();
    if (rz) *rz = cb.m_hitPointWorld.getZ();
    if (nx) *nx = cb.m_hitNormalWorld.getX();
    if (ny) *ny = cb.m_hitNormalWorld.getY();
    if (nz) *nz = cb.m_hitNormalWorld.getZ();
    if (modelId) *modelId = cb.m_collisionObject->getUserIndex();
    if (bodyTag) *bodyTag = cb.m_collisionObject->getUserIndex2();
    return 1;
}

int cm_raytest_ex(CMWorld w,
                  float sx, float sy, float sz,
                  float ex, float ey, float ez,
                  float *rx, float *ry, float *rz,
                  float *qx, float *qy, float *qz, float *qw,
                  float *px, float *py, float *pz,
                  int *modelId) {
    if (!w) return 0;
    auto *world = static_cast<CMWorldImpl*>(w);
    btVector3 Start(sx + 0.00001f, sy + 0.00001f, sz + 0.00001f);
    btVector3 End(ex, ey, ez);
    btCollisionWorld::ClosestRayResultCallback cb(Start, End);
    world->dynamicsWorld->rayTest(Start, End, cb);
    if (!cb.hasHit()) return 0;
    if (rx) *rx = cb.m_hitPointWorld.getX();
    if (ry) *ry = cb.m_hitPointWorld.getY();
    if (rz) *rz = cb.m_hitPointWorld.getZ();
    btQuaternion rot = cb.m_collisionObject->getWorldTransform().getRotation();
    btVector3    pos = cb.m_collisionObject->getWorldTransform().getOrigin();
    if (qx) *qx = rot.getX();
    if (qy) *qy = rot.getY();
    if (qz) *qz = rot.getZ();
    if (qw) *qw = rot.getW();
    if (px) *px = pos.getX();
    if (py) *py = pos.getY();
    if (pz) *pz = pos.getZ();
    if (modelId) *modelId = cb.m_collisionObject->getUserIndex();
    return 1;
}

int cm_raytest_all(CMWorld w,
                   float sx, float sy, float sz,
                   float ex, float ey, float ez,
                   float *rxArr, float *ryArr, float *rzArr,
                   float *distArr, int *modelArr, int *bodyTagArr,
                   int maxSize) {
    if (!w || maxSize <= 0) return -1;
    auto *world = static_cast<CMWorldImpl*>(w);
    btVector3 Start(sx + 0.00001f, sy + 0.00001f, sz + 0.00001f);
    btVector3 End(ex, ey, ez);
    btCollisionWorld::AllHitsRayResultCallback cb(Start, End);
    world->dynamicsWorld->rayTest(Start, End, cb);
    if (!cb.hasHit()) return 0;
    int count = cb.m_hitPointWorld.size();
    if (count > maxSize) return -1;
    for (int i = 0; i < count; i++) {
        if (rxArr)     rxArr[i]     = cb.m_hitPointWorld[i].getX();
        if (ryArr)     ryArr[i]     = cb.m_hitPointWorld[i].getY();
        if (rzArr)     rzArr[i]     = cb.m_hitPointWorld[i].getZ();
        if (modelArr)  modelArr[i]  = cb.m_collisionObjects[i]->getUserIndex();
        if (bodyTagArr) bodyTagArr[i] = cb.m_collisionObjects[i]->getUserIndex2();
        if (distArr) {
            float dx = cb.m_hitPointWorld[i].getX() - sx;
            float dy = cb.m_hitPointWorld[i].getY() - sy;
            float dz = cb.m_hitPointWorld[i].getZ() - sz;
            distArr[i] = sqrtf(dx*dx + dy*dy + dz*dz);
        }
    }
    return count;
}

int cm_contact_test(CMWorld w, CMShape s,
                     float x, float y, float z,
                     float qx, float qy, float qz, float qw) {
    if (!w || !s) return 0;
    auto *world = static_cast<CMWorldImpl*>(w);
    auto *shape = static_cast<CMShapeImpl*>(s);
    btDefaultMotionState *ms = new btDefaultMotionState(
        btTransform(btQuaternion(qx, qy, qz, qw), btVector3(x, y, z)));
    btRigidBody::btRigidBodyConstructionInfo ci(0, ms, shape->compound, btVector3(0, 0, 0));
    btRigidBody *body = new btRigidBody(ci);
    ContactSensor sensor;
    world->dynamicsWorld->contactTest(body, sensor);
    delete ms;
    delete body;
    return sensor.collided;
}

CMShape cm_load_dff(const char *dffPath, int modelid) {
	(void)dffPath;
    (void)modelid;
    return nullptr;
}

}
