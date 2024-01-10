'use client'

import './globals.css'
import { UserProvider } from "@/context/user";
import { LocalProvider } from "@/context/local";
import { Inter } from 'next/font/google'
import Header from "./header.js"
import Footer from "./footer.js"

const inter = Inter({ subsets: ['latin'] })
/*
export const metadata = {
  title: 'Thanet-PHA.org',
  description: "Thanet People's Health Alliance",
}
*/
export default function Layout({ children }) {
  return (
    <UserProvider>
        <LocalProvider>
          <Header/>
          {children}
          <Footer/>
        </LocalProvider>
    </UserProvider>
  )
}
