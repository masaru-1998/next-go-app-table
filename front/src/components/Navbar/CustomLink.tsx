import React, { useCallback } from 'react';
import { useRouter } from 'next/navigation';
import BottomNavigationAction, {
    BottomNavigationActionProps,
} from '@mui/material/BottomNavigationAction'
import { Button } from '@mui/material';;

interface CustomLinkProps extends Omit<BottomNavigationActionProps, 'component' | 'href' | 'onClick'> {
    to: string;
}

const CustomLink: React.FC<CustomLinkProps> = (props) => {
    const { to, ...other } = props;
    const router = useRouter();

    const handleClick = useCallback(
        (event: React.MouseEvent<HTMLButtonElement>) => {
            event.preventDefault();
            router.push(to);
        },
        [router, to],
    );

    return (
        <Button onClick={handleClick}>
            <BottomNavigationAction component="span" {...other} />
        </Button>
    );
};

CustomLink.displayName = 'CustomLink';

export default CustomLink;
