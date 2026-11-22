import { useAppSelector } from "@/store/hooks";
import CategoryCard from "./category-card";

const CategoryList = () => {
  const categories = useAppSelector((state) => state.category.categories);

  if (categories.length === 0) {
    return <p className="text-muted-foreground">Belum ada kategori.</p>;
  }

  return (
    <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
      {categories.map((category) => (
        <CategoryCard key={category.id} category={category} />
      ))}
    </div>
  );
};

export default CategoryList;
