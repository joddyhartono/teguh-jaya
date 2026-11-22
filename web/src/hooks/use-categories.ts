import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { fetchCategories } from "@/store/slices/category-slice";
import { useEffect } from "react";
import { toast } from "sonner";

const useCategories = () => {
  const dispatch = useAppDispatch();
  const categories = useAppSelector((state) => state.category.categories);
  const status = useAppSelector((state) => state.category.status);

  useEffect(() => {
    const load = async () => {
      try {
        await dispatch(fetchCategories()).unwrap();
      } catch (error) {
        if (typeof error === "object" && error !== null && "message" in error) {
          toast.error(String(error.message));
        }
      }
    };

    if (status === "idle") load();
  }, [status]);

  return {
    categories,
  };
};

export default useCategories;
