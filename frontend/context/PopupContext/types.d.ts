import { ReactNode } from 'react';

export type PopupType = '' | 'auth' | 'review';

export interface IPopupContext {
	popup: PopupType;
	setPopup: (value: PopupType) => void;
}

export interface IPopupProviderProps {
	children: ReactNode;
}
