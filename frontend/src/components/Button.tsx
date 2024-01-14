import React, { type ButtonHTMLAttributes, type ReactNode } from "react";
import "../styles/Button.css";

// Define props for the Button component
interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  // Add any additional custom props if needed
  children: ReactNode; // Make children required
}
const Button: React.FC<ButtonProps> = (props) => {
  const { children, ...rest } = props;

  return (
    <button className="button mt-5 py-2" onClick={props.onClick}>
      <span></span> <span></span> <span></span>
      <span></span>
      {children}
    </button>
  );
};

export default Button;
