import classes from  "./Footer.module.scss"

const Footer: React.FunctionComponent = (props) => {
    return (
        <footer className={classes.root}>
            <img src="img/logo_pix.png" alt="Ícone usuário" />
        </footer>
    );
};

export default Footer;