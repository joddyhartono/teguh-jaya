import { deleteCategory } from "@/services/category-service";
import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { fetchCategories } from "@/store/slices/category-slice";
import axios from "axios";
import { useEffect } from "react";
import { toast } from "sonner";

const useCategories = () => {
  const dispatch = useAppDispatch();
  const categories = useAppSelector((state) => state.category.categories);
  const status = useAppSelector((state) => state.category.status);

  {
    /* Error dari thunk itu hasil dari throw error baru dengan isinya makanya check di catch-nya seperti itu */
  }
  const load = async () => {
    try {
      await dispatch(fetchCategories()).unwrap();
    } catch (error) {
      if (typeof error === "object" && error !== null && "message" in error) {
        toast.error(String(error.message));
      }
    }
  };

  {
    /* Error dari async func biasa itu hasil dari axios error baru makanya check di catch-nya seperti itu */
  }
  const handleDelete = async (id: string) => {
    try {
      await deleteCategory(id);
      await load();
      toast.success("Kategori berhasil dihapus!");
    } catch (error) {
      if (axios.isAxiosError(error)) {
        toast.error(error.response?.data?.error ?? error.message);
      } else if (error instanceof Error) {
        toast.error(error.message);
      }
    }
  };

  useEffect(() => {
    if (status === "idle") load();
  }, [status]);

  return {
    categories,
    handleDelete,
  };
};

export default useCategories;
