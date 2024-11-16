import 'package:auth_repository/auth_repository.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_4_login/login/cubit/login_cubit.dart';
import 'package:flutter_4_login/login/screen/login_form.dart';

class LoginScreen extends StatelessWidget {
  const LoginScreen({super.key});

  static Route<void> route() {
    return MaterialPageRoute(
      builder: (_) => const LoginScreen(),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      // Latar belakang dengan gradien warna #CB9DF0 dan #FFF9BF
      body: Container(
        decoration: const BoxDecoration(
          gradient: LinearGradient(
            colors: [Color(0xFFCB9DF0), Color(0xFFFFF9BF)],
            begin: Alignment.topCenter,
            end: Alignment.bottomCenter,
          ),
        ),
        child: Center(
          child: Padding(
            padding: const EdgeInsets.all(24.0),
            child: BlocProvider<LoginCubit>(
              create: (context) {
                return LoginCubit(
                  authRepository: RepositoryProvider.of<AuthRepository>(context),
                );
              },
              // Kartu untuk formulir login
              child: Card(
                elevation: 12,
                shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(20),
                ),
                color: Colors.white, // Warna dasar kartu
                child: Padding(
                  padding: const EdgeInsets.all(20.0),
                  child: Column(
                    mainAxisSize: MainAxisSize.min,
                    crossAxisAlignment: CrossAxisAlignment.stretch,
                    children: [
                      const Text(
                        'LOGIN',
                        textAlign: TextAlign.center,
                        style: TextStyle(
                          fontSize: 28,
                          fontWeight: FontWeight.bold,
                          color: Color(0xFF6A1B9A), // Ungu gelap untuk aksen
                        ),
                      ),
                      const SizedBox(height: 16),
                      // Tambahan ikon login di bawah teks
                      Icon(
                        Icons.login_rounded,
                        size: 64,
                        color: Color(0xFF6A1B9A), // Warna ungu serupa
                      ),
                      const SizedBox(height: 20),
                      // Formulir Login
                      const LoginForm(),
                      const SizedBox(height: 20),
                      // Tambahan footer kecil
                     
                    ],
                  ),
                ),
              ),
            ),
          ),
        ),
      ),
    );
  }
}
