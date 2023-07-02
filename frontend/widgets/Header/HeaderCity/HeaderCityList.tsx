import { cities } from '@/constants';
import { ICity } from '@/constants/types'
import Image from 'next/image';
import { ChangeEvent, useState } from 'react';
import styles from '../styles.module.scss';

interface IHeaderCityListProps {
	activeId?: number;
	isVisible: boolean;
	updateActiveId: (id: number) => void;
	hideList: () => void;
}

const HeaderCityList: React.FC<IHeaderCityListProps> = ({
	activeId,
	isVisible,
	updateActiveId,
	hideList,
}) => {
	const [query, setQuery] = useState<string>('');
	const filteredCities = cities?.filter(obj =>
		obj.name.toLocaleLowerCase().includes(query.trim().toLowerCase())
	);

	const updateQuery = (event: ChangeEvent<HTMLInputElement>): void =>
		setQuery(event.target.value);
	
	const cityClickHandler = (city: ICity) => {
		updateActiveId(city.id);
		setQuery('');
		hideList();
	};	

	return (
		<div
			className={`
				absolute lg:bg-[#84EA94] lg:rounded-2xl rounded-b-2xl py-3 left-1/2 -translate-x-1/2 min-w-[230px]
				lg:shadow-[0px_5px_10px_rgba(0,0,0,.25)] shadow-[0px_12px_10px_rgba(0,0,0,.2)] transition-all duration-200 w-full
				${
					isVisible
						? 'opacity-100 visible lg:top-[calc(100%_+_10px)] top-full'
						: 'opacity-0 invisible top-[calc(100%_+_20px)]'
				}
				${styles['city-list']}
			`}
		>
			<div className='lg:px-3 px-4 max-w-full flex flex-row justify-between items-center gap-3 pb-3 border-b-[2.5px] border-[#414141]'>
				<Image src='/icons/search.svg' width={20} height={24} alt='' />
				<input
					className='w-full text-[#414141] placeholder:text-inherit font-medium lg:text-[18px] md:text-[22px] text-[16px] leading-none bg-transparent'
					type='text'
					placeholder='Поиск'
					value={query}
					onChange={updateQuery}
				/>
				<Image
					className='cursor-pointer transition-opacity hover:opacity-70 duration-200'
					src='/icons/close.svg'
					width={20}
					height={20}
					alt=''
					onClick={hideList}
				/>
			</div>
			<ul className='mt-3 space-y-1.5 lg:max-h-[230px] md:max-h-[313px] max-h-[212px] overflow-y-auto'>
				{filteredCities.length ? (
					filteredCities.map(obj => (
						<li
							onClick={() => cityClickHandler(obj)}
							key={obj.id}
							className={`
								cursor-pointer lg:px-3 px-4 lg:text-[18px] md:text-[22px] text-[16px] pb-1.5 last:pb-0 border-b-[2px] text-[#414141] border-[#414141] last:border-b-0
								transition-colors duration-300 hover:font-bold hover:text-black hover:border-black
								${activeId === obj.id ? 'font-bold text-black border-black' : ''}
							`}
						>
							{obj.name}
						</li>
					))
				) : (
					<li className='px-3 lg:text-[18px] md:text-[22px] text-[16px] text-[#414141]'>Ничего не найдено</li>
				)}
			</ul>
		</div>
	);
};

export default HeaderCityList;
