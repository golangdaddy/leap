'use client'

import './globals.css'
import { UserProvider } from "@/context/user";
import { LocalProvider } from "@/context/local";
import { MessagingProvider } from "@/context/messaging";
import { Inter } from 'next/font/google'
import Header from "../components/header.js"
import Footer from "../components/footer.js"

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
						<span className='text-black'>{children}</span>
						<Footer/>
					</MessagingProvider>
				</LocalProvider>
		</UserProvider>
	)
}
