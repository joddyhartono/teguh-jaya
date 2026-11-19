import type React from "react";
import { Label } from "../ui/label";
import { Input } from "../ui/input";

type FormFieldProps = React.ComponentProps<"input"> & {
  label: string;
  name: string;
  error?: string;
};

const FormField = ({ label, name, error, ...props }: FormFieldProps) => {
  return (
    <div className="w-full flex flex-col gap-2 justify-between">
      <Label htmlFor={name} className={error ? "text-red-500" : ""}>
        {label}
      </Label>
      <Input
        id={name}
        name={name}
        {...props}
        className={error ? "border-red-500 text-red-500" : ""}
      />
      {error && <p className="text-red-500">{error}</p>}
    </div>
  );
};

export default FormField;
