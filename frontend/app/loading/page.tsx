import SearchWrapper from '@/shared/SearchWrapper';
import Loading from '@/components/Loading';

export const metadata = {
	title: 'Загрузка...',
};

const LoadingPage: React.FC = () => {
	return (
		<section className='section'>
			<div className='container-md container-full md:pt-[129px] pt-[115px]'>
				<SearchWrapper>
					<Loading />
				</SearchWrapper>
			</div>
		</section>
	);
};

export default LoadingPage;
