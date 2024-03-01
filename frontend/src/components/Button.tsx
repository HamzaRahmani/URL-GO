import React, { type ButtonHTMLAttributes, type ReactNode } from "react";
import "../styles/Button.css";

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  children: ReactNode;
}
const Button: React.FC<ButtonProps> = (props) => {
  const { children, ...rest } = props;

  return (
    <button className="button mt-5 py-3 mx-1 w-5/6" onClick={props.onClick}>
      <span></span> <span></span> <span></span>
      <span></span>
      {children}
    </button>
  );
};

export default Button;
