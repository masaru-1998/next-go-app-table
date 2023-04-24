'use client';
import './globals.css'
import { SideNavbar } from '@/components/Navbar';
import styles from './layout.module.css';
import Head from 'next/head'

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="ja">
      <Head>
        <title>Shcedule Generator</title>
      </Head>
      <body className={styles.body}>
        <div className={styles.main}>
          <SideNavbar />
          {children}
        </div>
      </body>
    </html>
  )
}
