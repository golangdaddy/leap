'use client'

import './globals.css'
import { UserProvider } from "@/context/user";
import { LocalProvider } from "@/context/local";
import { MessagingProvider } from "@/context/messaging";
import { Inter } from 'next/font/google'
import Header from "./header.js"
import Footer from "./footer.js"

const inter = Inter({ subsets: ['latin'] })
/*
export const metadata = {}
*/
export default function Layout({ children }) {
	return (
		<UserProvider>
				<LocalProvider>
					<MessagingProvider>
						<Header/>
						{children}
						<Footer/>
					</MessagingProvider>
				</LocalProvider>
		</UserProvider>
	)
}
