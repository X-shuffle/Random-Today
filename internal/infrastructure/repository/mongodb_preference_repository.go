package repository

import (
	"context"
	"errors"
	"math/rand"
	"random_today/internal/domain/entity"
	"random_today/internal/infrastructure/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBPreferenceRepository MongoDB偏好仓储实现
type MongoDBPreferenceRepository struct {
	collection *mongo.Collection
}

// NewMongoDBPreferenceRepository 创建MongoDB偏好仓储实例
func NewMongoDBPreferenceRepository() *MongoDBPreferenceRepository {
	db := database.GetDatabase()
	collection := db.Collection("preferences")
	return &MongoDBPreferenceRepository{
		collection: collection,
	}
}

// Create 创建偏好
func (r *MongoDBPreferenceRepository) Create(ctx context.Context, preference *entity.Preference) error {
	// 生成ObjectID
	objectID := primitive.NewObjectID()
	preference.ID = objectID.Hex()

	// 插入文档
	_, err := r.collection.InsertOne(ctx, preference)
	return err
}

// GetByID 根据ID获取偏好
func (r *MongoDBPreferenceRepository) GetByID(ctx context.Context, id string) (*entity.Preference, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var preference entity.Preference
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&preference)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("preference not found")
		}
		return nil, err
	}

	return &preference, nil
}

// GetByType 根据类型获取偏好列表
func (r *MongoDBPreferenceRepository) GetByType(ctx context.Context, preferenceType entity.PreferenceType) ([]*entity.Preference, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"type": preferenceType})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var preferences []*entity.Preference
	if err = cursor.All(ctx, &preferences); err != nil {
		return nil, err
	}

	return preferences, nil
}

// GetAll 获取所有偏好
func (r *MongoDBPreferenceRepository) GetAll(ctx context.Context) ([]*entity.Preference, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var preferences []*entity.Preference
	if err = cursor.All(ctx, &preferences); err != nil {
		return nil, err
	}

	return preferences, nil
}

// Update 更新偏好
func (r *MongoDBPreferenceRepository) Update(ctx context.Context, preference *entity.Preference) error {
	objectID, err := primitive.ObjectIDFromHex(preference.ID)
	if err != nil {
		return err
	}

	preference.UpdatedAt = time.Now()
	_, err = r.collection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": preference},
	)
	return err
}

// Delete 删除偏好
func (r *MongoDBPreferenceRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

// DeleteByType 根据类型删除偏好
func (r *MongoDBPreferenceRepository) DeleteByType(ctx context.Context, preferenceType entity.PreferenceType) error {
	_, err := r.collection.DeleteMany(ctx, bson.M{"type": preferenceType})
	return err
}

// CountByType 根据类型获取文档总数
func (r *MongoDBPreferenceRepository) CountByType(ctx context.Context, preferenceType entity.PreferenceType) (int64, error) {
	filter := bson.M{"type": preferenceType}
	return r.collection.CountDocuments(ctx, filter)
}

// GetByTypeAndIndex 根据类型和索引位置获取偏好
func (r *MongoDBPreferenceRepository) GetByTypeAndIndex(ctx context.Context, preferenceType entity.PreferenceType, index int64) (*entity.Preference, error) {
	filter := bson.M{"type": preferenceType}

	// 使用skip来获取指定索引位置的文档
	opts := options.FindOne().SetSkip(index)

	var preference entity.Preference
	err := r.collection.FindOne(ctx, filter, opts).Decode(&preference)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &preference, nil
}

// GetRandomByType 根据类型随机获取一个偏好（优化版本）
func (r *MongoDBPreferenceRepository) GetRandomByType(ctx context.Context, preferenceType entity.PreferenceType) (*entity.Preference, error) {
	// 获取该类型的文档总数
	total, err := r.CountByType(ctx, preferenceType)
	if err != nil {
		return nil, err
	}

	if total == 0 {
		return nil, errors.New("no preferences found for this type")
	}

	// 随机选择一个索引
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Int63n(total)

	// 通过索引获取特定文档
	return r.GetByTypeAndIndex(ctx, preferenceType, randomIndex)
}
