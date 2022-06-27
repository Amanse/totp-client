import { Writable, writable } from "svelte/store";
import type { Response } from "./Types";

export const AllData: Writable<Partial<Response>>= writable({"issuers": {}})