import { serializeStyles } from '@emotion/serialize';
import React from 'react';

const sizes = {

}

type ButtonProps = React.ButtonHTMLAttributes<HTMLButtonElement>& {
    children: React.ReactNode;
}
export const Button = ({ children, ...props }: ButtonProps) => {
    return (
        <button {...props}>
            {children}
        </button>
    )
}
