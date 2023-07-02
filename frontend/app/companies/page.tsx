import SearchWrapper from '@/shared/SearchWrapper';
import Companies from '@/components/Companies';

export const metadata = {
	title: 'Результат',
};

const CompaniesPage: React.FC = () => {
	return (
		<section className='section'>
			<div className='container-md container-full md:pt-[129px] pt-[115px]'>
				<SearchWrapper>
					<Companies />
				</SearchWrapper>
			</div>
		</section>
	);
};

export default CompaniesPage;
