'use client';

import { createContext, useEffect, useState } from 'react';
import { IPopupContext, IPopupProviderProps, PopupType } from './types';
import { handleScrollBarPadding } from '@/utils/helpers';

const DEFAULT_POPUP_CONTEXT: IPopupContext = {
	popup: '',
	setPopup: () => {},
};

export const PopupContext = createContext<IPopupContext>(DEFAULT_POPUP_CONTEXT);

const PopupProvider: React.FC<IPopupProviderProps> = ({ children }) => {
	const [popup, setPopup] = useState<PopupType>('');

	const popupHandler = (value: PopupType): void => setPopup(value);

	useEffect(() => {
		let isMounted: boolean = true;
		const documentWidth: number = document.documentElement.clientWidth;
		const windowWidth: number = window.innerWidth;
		const scrollBarWidth: number = windowWidth - documentWidth;
		const layout: HTMLElement | null = document.querySelector('.layout');
		const fixedElements: NodeListOf<HTMLElement> = document.querySelectorAll('.lock-padding');
		
		if(!layout) return;

		if(isMounted && popup.length) {
			document.body.classList.add('lock');
			handleScrollBarPadding(scrollBarWidth, layout, ...fixedElements);	
		} else {
			setTimeout(() => {
				document.body.classList.remove('lock');
				handleScrollBarPadding(0, layout, ...fixedElements);
			}, 400);
		}
		
		return () => { isMounted = false };
	}, [popup]);

	return (
		<PopupContext.Provider
			value={{
				popup: popup,
				setPopup: popupHandler,
			}}
		>
			{children}
		</PopupContext.Provider>
	);
};

export default PopupProvider;
