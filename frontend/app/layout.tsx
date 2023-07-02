import PopupProvider from '@/context/PopupContext';
import Footer from '@/widgets/Footer';
import Header from '@/widgets/Header';
import Pattern from '@/widgets/Pattern';
import Popups from '@/widgets/Popups';
import { Inter } from 'next/font/google';
import './globals.scss';
import { ReactNode } from 'react';

const inter = Inter({
	subsets: ['cyrillic'],
	weight: ['400', '500', '600', '700', '800'],
});

const RootLayout = ({ children }: { children: ReactNode }) => {
	return (
		<html lang='ru'>
			<body className={inter.className}>
				<PopupProvider>
					<div className='layout'>
						<Pattern />
						<Header />
						<main>{children}</main>
						<Footer />
						<Popups />
					</div>
				</PopupProvider>
			</body>
		</html>
	);
};

export default RootLayout;
