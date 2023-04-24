'use client';
import React from 'react';
import Link from 'next/link';
import Box from '@mui/material/Box';
import BottomNavigation from '@mui/material/BottomNavigation';

import styles from './styles/SideNavbar.module.css'
import CustomLink from './CustomLink';

export const SideNavbar = () => {
    const [value, setValue] = React.useState('');
    return (
        <Box sx={{
            minWidth: 250,
            height: '100vh !important',
            borderRight: '2px solid #000',
            justifyContent: 'center'
        }}>
            <h1 className={styles.sideTitle}>
                <Link
                    className={styles.sideTitleLink}
                    href="/"
                >Back Offiece</Link>
            </h1>
            <BottomNavigation
                sx={{flexDirection: 'column', justifyContent: 'center'}}
                showLabels
                value={value}
                onChange={(event, newValue) => {
                    setValue(newValue);
                }}
            >
                <CustomLink
                    to="/report"
                    label="報告"
                    classes={{ root: styles.sideNavbarButtom }}
                />
                <CustomLink
                    to="/schedule"
                    label="スケジュール作成"
                    classes={{ root: styles.sideNavbarButtom }}
                />
                <CustomLink
                    to="/manage-task"
                    label="タスク管理"
                    classes={{ root: styles.sideNavbarButtom }}
                />
            </BottomNavigation>
        </Box>
    );
}
