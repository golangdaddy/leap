'use client'

import * as React from 'react'
import { useState } from 'react'

import Link from 'next/link'

export default function Footer() {

    const [contacts, setContacts] = useState(false)
    const [disclaimer, setDisclaimer] = useState(false)

    function updateSearch() {
        console.log(query)
    }

    function toggleContacts() {
        setContacts(!contacts)
        setDisclaimer(false)
    }
    function toggleDisclaimer() {
        setDisclaimer(!disclaimer);
        setContacts(false)
    }

    return (
    <footer>
        

<footer className="bg-white rounded-lg shadow m-4 dark:bg-gray-800">
    <div className="w-full mx-auto max-w-screen-xl p-4 md:flex md:items-center md:justify-between">
      <span className="text-sm text-gray-500 sm:text-center dark:text-gray-400">© 2024 <a href="https://{{.WebsiteName}}.com/" className="hover:underline">{{.WebsiteName}}</a>. All Rights Reserved.
    </span>
    <ul className="flex flex-wrap items-center mt-3 text-sm font-medium text-gray-500 dark:text-gray-400 sm:mt-0">
        <li>
            <a href="#" className="hover:underline me-4 md:me-6">About</a>
        </li>
        <li>
            <a href="#" className="hover:underline me-4 md:me-6">Policy</a>
        </li>
        <li>
            <a href="#" className="hover:underline me-4 md:me-6">Licensing</a>
        </li>
        <li>
            <a href="#" className="hover:underline">Contact</a>
        </li>
    </ul>
    </div>
</footer>

    </footer>
    )
}