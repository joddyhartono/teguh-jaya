import useCategories from "@/hooks/use-categories";
import { useAppSelector } from "@/store/hooks";
import CategoryList from "./category-list";

const CategorySection = () => {
  const status = useAppSelector((state) => state.category.status);
  const { categories } = useCategories();

  return (
    <section className="flex flex-col gap-4 p-4">
      <h2 className="text-xl font-semibold">Kategori</h2>

      {status === "loading" && <p>Memuat...</p>}
      {status === "succeeded" && <CategoryList />}
    </section>
  );
};

export default CategorySection;
