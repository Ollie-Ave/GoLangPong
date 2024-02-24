package entities

type IEntityManager interface {
    GetEntities() []IEntity
    SpawnEntity(IEntity)
}

type EntityManager struct {
    entities []IEntity
}

func NewEntityManager() *EntityManager {
    return &EntityManager {}
}

func (entityManager *EntityManager) SpawnEntity(entity IEntity) {
    entityManager.entities = append(entityManager.entities, entity)
}

func (entityManager *EntityManager) GetEntities() []IEntity {
    return entityManager.entities
}
