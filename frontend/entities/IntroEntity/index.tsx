import Title from '@/shared/Title';
import Greeting from '@/components/Greeting';
import Calculator from '@/components/Calculator';

const IntroEntity: React.FC = () => {
	return (
		<section className='section'>
			<div className='container-md'>
				<Title type='page'>
					Займ <span className='text-tertiary'>онлайн</span> в кратчайшие сроки
				</Title>
			</div>
			<div className='container-sm'>
				<Greeting />
				<Calculator />
			</div>
		</section>
	);
};

export default IntroEntity;
