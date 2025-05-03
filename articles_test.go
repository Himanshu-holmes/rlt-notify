package main

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/himanshu-holmes/rlt-notify/repository"
	"github.com/stretchr/testify/assert"
)

type mockArticle struct {
	items map[uint64]repository.Article
}

func (m *mockArticle) ById(ctx context.Context, id int) (repository.Article, error) {
   val, has := m.items[uint64(id)]
   if !has {
      return repository.Article{}, repository.ErrNotFound
	}
	return val, nil
}

func TestArticle_ByID(t *testing.T) {
   ma := &mockArticle{items: map[uint64]repository.Article{
      1: {
         ID:      1,
         Title:   "Title#1",
         Content: "content of the first article.",
      },
   }}
   a := NewArticle(ma, 3)

   _, err := a.ById(context.Background(), 10)
   assert.ErrorIs(t, repository.ErrNotFound, err)

   item, err := a.ById(context.Background(), 1)
   fmt.Printf("%+v\n", item)
   // print item.summary
   fmt.Printf("%+v\n", item.Summary)
   assert.NoError(t, err)
   assert.Equal(t, "Title#1", item.Title)
   assert.Equal(t, "https://site.com/a/1", item.More)
   assert.Equal(t, uint64(1), item.ID)
   assert.Equal(t, "content of the", item.Summary)
}

func BenchmarkArticle(b *testing.B) {
   ma := &mockArticle{items: map[uint64]repository.Article{
      1: {
         ID:      1,
         Title:   "Title#1",
         Content: "content of the first article.",
      },
   }}
   a := NewArticle(ma, 3)

   for i := 0; i < b.N; i++ {
      a.ById(context.Background(), 10)
   }
}


// CheckEveryItem looks for the given lookup argument in the slice and returns its index if it is presented.
// Otherwise, it returns -1.
func CheckEveryItem(items []int, lookup int) int {
 for i := 0; i < len(items); i++ {
  if items[i] == lookup {
   return i
  }
 }
 return -1
}

// BinarySearch expects to receive a sorted slice and looks for the index of the given value accordingly.
func BinarySearch(items []int, lookup int) int {
 left := 0
 right := len(items) - 1
 for {
  if left == lookup {
   return left
  }
  if right == lookup {
   return right
  }

  center := (right + left) / 2
  if items[center] == lookup {
   return center
  }
  if center > lookup {
   right = center
  }
  if center < lookup {
   left = center
  }
  if left >= right-1 {
   return -1
  }
 }
}

type Algorithm func(items []int, lookup int) int

func testAlgorithm(alg Algorithm, t *testing.T) {
   items := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
   for i := 0; i <= 9; i++ {
      assert.Equal(t, i, alg(items, i))
   }

   assert.Equal(t, -1, alg(items, 100))
}

func TestCheckEveryItem(t *testing.T) {
   testAlgorithm(CheckEveryItem, t)
}

func TestBinarySearch(t *testing.T) {
   testAlgorithm(BinarySearch, t)
}

func benchmarkAlgorithm(alg Algorithm, b *testing.B) {
   totalItems := int(1e12)
   items := make([]int, totalItems)
   for i := 0; i < totalItems; i++ {
      items[i] = i
   }
   b.ResetTimer()
   for i := 0; i < b.N; i++ {
      lookup := rand.Intn(totalItems - 1)
      alg(items, lookup)
   }
}

func BenchmarkCheckEveryItem(b *testing.B) {
   benchmarkAlgorithm(CheckEveryItem, b)
}
func BenchmarkBinarySearch(b *testing.B) {
   benchmarkAlgorithm(BinarySearch, b)
}