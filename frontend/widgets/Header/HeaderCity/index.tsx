import { cities } from '@/constants';
import { ICity } from '@/constants/types';
import Image from 'next/image';
import React, { useState } from 'react';
import HeaderCityList from './HeaderCityList';

const HeaderCity: React.FC = () => {
	const [activeId, setActiveId] = useState<number>(1);
	const [isVisible, setIsVisible] = useState<boolean>(false);

	const activeCity: ICity | undefined = cities.find(obj => obj.id === activeId);

	const updateActiveId = (id: number): void => setActiveId(id);
	const toggleList = (): void => setIsVisible(prev => !prev);
	const hideList = (): void => setIsVisible(false);

	return (
		<div className='lg:relative z-[2]'>
			<div
				onClick={toggleList}
				className='text-center underline 2xl:text-[20px] text-[17px] cursor-pointer leading-none transition-opacity duration-200 hover:opacity-70'
			>
				<Image
					className='mx-auto 2xl:h-[50px] lg:h-[40px] md:h-[60px] h-[40px]'
					src='/icons/location.svg'
					width={44}
					height={50}
					alt=''
				/>
				<div className='mt-2 lg:block hidden'>{activeCity?.name}</div>
			</div>
			<HeaderCityList
				activeId={activeId}
				updateActiveId={updateActiveId}
				isVisible={isVisible}
				hideList={hideList}
			/>
		</div>
	);
};

export default HeaderCity;
