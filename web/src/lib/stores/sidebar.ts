import { createLocalStorage, persist } from '@macfja/svelte-persistent-store';
import { writable } from 'svelte/store';

export const sidebar = persist(writable(false), createLocalStorage(), 'sidebar');
