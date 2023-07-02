import ReviewEntity from '@/entities/ReviewEntity';
import Title from '@/shared/Title';

export const metadata = {
	title: 'Отзывы',
};

const ReviewsPage = () => {
	return (
		<section className='section'>
			<div className='container-xl'>
				<Title type='page'>
					Мнения о <span className='text-primary'>Займике</span>
				</Title>
				<p className='lg:text-[30px] md:text-[26px] text-[18px] leading-tight font-medium lg:mt-4 mt-2'>
					Если у вас есть сомнения по поводу предостовляемых нами услуг, то
					можете ознакомиться с опытом получения займов других людей и почитать
					их отзывы о нашей работе ниже. Уверен, вы сделаете правильный выбор.
				</p>
				<ReviewEntity />
			</div>
		</section>
	);
};

export default ReviewsPage;
