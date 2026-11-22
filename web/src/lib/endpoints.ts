const AUTH = {
  LOGIN: "/auth/login",
} as const;

const PRODUCTS = {
  BASE: "/products",
  DETAIL: (id: string) => `products/${id}`,
  BY_CATEGORY: (categoryId: string) => `products?categoryId=${categoryId}`,
} as const;

const CATEGORIES = {
  BASE: "/categories",
} as const;

export { AUTH, PRODUCTS, CATEGORIES };
