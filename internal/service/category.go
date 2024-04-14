package service

import (
	"context"
	"fc3-grpc/internal/database"
	"fc3-grpc/internal/pb"
	"io"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (categoryService *CategoryService) CreateCategory(ctx context.Context, input *pb.CreateCategoryRequest) (*pb.Category, error) {
	newCategory, err := categoryService.CategoryDB.Create(input.Name, input.Description)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create category: %v", err)
	}
	category := &pb.Category{
		Id:          newCategory.ID,
		Name:        newCategory.Name,
		Description: newCategory.Description,
	}
	return category, nil
}

func (categoryService *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryList{}
	for {
		input, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}
		if err != nil {
			return status.Errorf(codes.Internal, "Failed to receive category: %v", err)
		}
		newCategory, err := categoryService.CategoryDB.Create(input.Name, input.Description)
		if err != nil {
			return status.Errorf(codes.Internal, "Failed to create category: %v", err)
		}
		category := &pb.Category{
			Id:          newCategory.ID,
			Name:        newCategory.Name,
			Description: newCategory.Description,
		}
		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          category.Id,
			Name:        category.Name,
			Description: category.Description,
		})
	}
}

func (categoryService *CategoryService) CreateCategoryStreamBidirectional(stream pb.CategoryService_CreateCategoryStreamBidirectionalServer) error {
	for {
		input, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return status.Errorf(codes.Internal, "Failed to receive category: %v", err)
		}
		newCategory, err := categoryService.CategoryDB.Create(input.Name, input.Description)
		if err != nil {
			return status.Errorf(codes.Internal, "Failed to create category: %v", err)
		}
		category := &pb.Category{
			Id:          newCategory.ID,
			Name:        newCategory.Name,
			Description: newCategory.Description,
		}
		if err := stream.Send(category); err != nil {
			return status.Errorf(codes.Internal, "Failed to send category: %v", err)
		}
	}
}

func (categoryService *CategoryService) ListCategories(ctx context.Context, _ *pb.Blank) (*pb.CategoryList, error) {
	categories, err := categoryService.CategoryDB.FindAll()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to list categories: %v", err)
	}
	var pbCategories []*pb.Category
	for _, category := range categories {
		pbCategory := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}
		pbCategories = append(pbCategories, pbCategory)
	}
	return &pb.CategoryList{Categories: pbCategories}, nil
}
