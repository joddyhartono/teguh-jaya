import { Link } from "react-router";
import { Button } from "../ui/button";
import useAuth from "@/hooks/use-auth";

const Navbar = () => {
  const { logout } = useAuth();

  return (
    <nav className="flex justify-between items-center w-full p-4 border-b bg-primary text-secondary">
      <Link to="/admin/dashboard">
        <img src="/logo.png" alt="logo" className="h-12 w-auto" />
      </Link>
      {/* <ul className="flex items-center gap-8">
        <li>
          <Link to="/admin/categories">Kategori</Link>
        </li>
        <li>
          <Link to="/admin/products">Produk</Link>
        </li>
      </ul> */}
      <Button className="cursor-pointer" onClick={logout}>
        Logout
      </Button>
    </nav>
  );
};

export default Navbar;
