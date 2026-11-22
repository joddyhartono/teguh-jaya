import FormField from "@/components/shared/form-field";
import { Button } from "@/components/ui/button";
import useAuth from "@/hooks/use-auth";
import { useAppSelector } from "@/store/hooks";
import React, { useState } from "react";
import { Navigate, useNavigate } from "react-router";
import { toast } from "sonner";

const LoginPage = () => {
  const token = useAppSelector((state) => state.auth.token);
  if (token) {
    <Navigate to="/admin/dashboard" replace={true} />;
  }

  const { login } = useAuth();

  const [form, setForm] = useState({
    username: "",
    password: "",
  });

  const [errors, setErrors] = useState({
    username: "",
    password: "",
  });

  const [loading, setLoading] = useState(false);

  const navigate = useNavigate();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm((prev) => ({ ...prev, [e.target.name]: e.target.value }));
    setErrors((prev) => ({ ...prev, [e.target.name]: "" }));
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (form.username === "") {
      setErrors((prev) => ({
        ...prev,
        username: "Username wajib diisi",
      }));
    }

    if (form.password === "") {
      setErrors((prev) => ({
        ...prev,
        password: "Password wajib diisi",
      }));
    }

    if (form.username === "" || form.password === "") {
      return;
    }

    setLoading(true);
    try {
      await login(form);
      navigate("/admin/dashboard", { replace: true });
    } catch (error) {
      if (typeof error === "object" && error !== null && "message" in error) {
        toast.error(String(error.message));
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    <section className="flex flex-col gap-8 justify-center items-center min-h-screen mx-auto">
      <h1 className="text-xl font-semibold">Teguh Jaya</h1>
      <form
        className="flex flex-col gap-8 justify-between items-start max-w-sm w-full p-8 border rounded-md"
        onSubmit={handleSubmit}
      >
        <h2 className="text-lg font-medium">Sign in to your account</h2>
        <FormField
          label="Username"
          name="username"
          onChange={handleChange}
          error={errors.username}
        />
        <FormField
          label="Password"
          name="password"
          type="password"
          onChange={handleChange}
          error={errors.password}
        />
        <Button type="submit" className="w-full" disabled={loading}>
          Sign in
        </Button>
      </form>
    </section>
  );
};

export default LoginPage;
