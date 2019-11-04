import { writable } from "svelte/store";

// User contains information about the current user
export const user = writable({});

export const sidebarVisible = writable({});