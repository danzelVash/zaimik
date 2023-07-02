import Raccoon from '@/shared/Raccoon';
import SearchText from '@/shared/SearchText';
import SearchTitle from '@/shared/SearchTitle';
import Link from 'next/link';
import CompaniesItems from './CompaniesItems';

const Companies: React.FC = () => {
	return (
		<div>
			<Raccoon
				src='/static/raccoons/money_raccoon.png'
				width={213}
				height={258}
				className='md:w-auto md:h-auto w-[190px] h-[230px] absolute z-[-1] left-1/2 top-0 -translate-x-1/2 -translate-y-1/2'
			/>
			<SearchTitle>
				Миссия выполнена! <br />Я подобрал вам лучшие варианты займов!
			</SearchTitle>
			<SearchText className='lg:mt-4 mt-3'>
				Чтобы точно получить микрозайм, подайте заявку в несколько
				микрофинансовых объединений
			</SearchText>
			<CompaniesItems />
			<div className='text-center lg:mt-7 md:mt-6 mt-5'>
				<Link
					className='lg:text-[28px] md:text-[24px] text-[18px] underline transition-opacity duration-200 hover:opacity-70'
					href='/'
				>
					Вернуть на главную
				</Link>
			</div>
		</div>
	);
};

export default Companies;
