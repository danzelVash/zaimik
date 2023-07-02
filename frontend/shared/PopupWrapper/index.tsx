'use client';

import { PopupContext } from '@/context/PopupContext';
import { PopupType } from '@/context/PopupContext/types';
import Image from 'next/image';
import { ReactNode, useContext } from 'react';

interface IPopupWrapperProps {
	children: ReactNode;
	currentPopup: PopupType;
	className?: string;
}

const PopupWrapper: React.FC<IPopupWrapperProps> = ({
	children,
	currentPopup,
	className = '',
}) => {
	const { popup, setPopup } = useContext(PopupContext);
	const isVisible = currentPopup === popup;

	return (
		<div
			className={`${
				isVisible ? 'opacity-100 visible' : 'invisible opacity-0'
			} transition-all duration-[400ms] fixed top-0 left-0 w-full h-full z-[20] overflow-y-auto`}
		>
			<div className='w-full min-h-full flex items-center justify-center bg-[rgba(0,0,0,.5)] px-4 py-8'>
				<div
					className={`${className} ${
						isVisible ? '' : 'translate-y-[30px]'
					} transition-all duration-[400ms] w-full main-gradient rounded-3xl relative xl:py-[60px] md:py-14 py-10 xl:px-[120px] md:px-16 px-5`}
				>
					{children}
					<Image
						className='absolute top-0 left-0 w-full h-full z-[1] object-cover'
						src='/static/popups/pattern.png'
						fill
						alt=''
					/>
					<div
						onClick={() => setPopup('')}
						className='cursor-pointer transition-opacity duration-200 hover:opacity-70 absolute md:top-5 top-3.5 md:right-5 right-3.5 z-[2]'
					>
						<Image
							className='md:w-[30px] md:h-[30px] w-[20px] h-[20px]'
							src='/icons/close.svg'
							width={30}
							height={30}
							alt=''
						/>
					</div>
				</div>
			</div>
			<div
				onClick={() => setPopup('')}
				className='absolute top-0 left-0 w-full h-full'
			></div>
		</div>
	);
};

export default PopupWrapper;
